package indexhandler

import (
	"html/template"
	"net/http"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/skymarshal/token"
	"github.com/gobuffalo/packr"
)

type templateData struct {
	CSRFToken string
	AuthToken string
}

type handler struct {
	logger     lager.Logger
	template   *template.Template
	middleware token.Middleware
}

func NewHandler(logger lager.Logger, middleware token.Middleware) (http.Handler, error) {
	tfuncs := &templateFuncs{
		assetIDs: map[string]string{},
	}

	funcs := template.FuncMap{
		"asset": tfuncs.asset,
	}

	box := packr.NewBox("../public")

	src, err := box.MustBytes("index.html")
	if err != nil {
		return nil, err
	}

	t, err := template.New("index").Funcs(funcs).Parse(string(src))
	if err != nil {
		return nil, err
	}

	return &handler{
		logger:     logger,
		template:   t,
		middleware: middleware,
	}, nil
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.logger.Session("index")

	authToken := h.middleware.GetAuthToken(r)
	csrfToken := h.middleware.GetCSRFToken(r)

	err := h.template.Execute(w, templateData{
		CSRFToken: csrfToken,
		AuthToken: authToken,
	})

	if err != nil {
		log.Fatal("failed-to-build-template", err, lager.Data{})
		w.WriteHeader(http.StatusInternalServerError)
	}
}
