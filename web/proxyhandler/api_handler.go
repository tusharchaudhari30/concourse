package proxyhandler

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/skymarshal/token"
)

const CSRFHeaderName = "X-Csrf-Token"

var ErrStatusUnauthorized = errors.New("401 Unauthorized")

func NewApiHandler(
	logger lager.Logger,
	target *url.URL,
	xFrameOptions string,
	middleware token.Middleware,
) *apiHandler {

	dialer := &net.Dialer{
		Timeout:   24 * time.Hour,
		KeepAlive: 24 * time.Hour,
	}

	transport := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		Dial:                dialer.Dial,
		TLSHandshakeTimeout: 60 * time.Second,
	}

	handler := httputil.NewSingleHostReverseProxy(target)
	handler.FlushInterval = 100 * time.Millisecond
	handler.Transport = transport

	handler.ModifyResponse = func(r *http.Response) error {
		switch r.StatusCode {
		case http.StatusUnauthorized:
			return ErrStatusUnauthorized
		default:
			return nil
		}
	}

	handler.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		switch err {
		case ErrStatusUnauthorized:
			middleware.UnsetCSRFToken(w)
			middleware.UnsetAuthToken(w)
			w.WriteHeader(http.StatusUnauthorized)
		default:
			w.WriteHeader(http.StatusBadGateway)
		}
	}

	return &apiHandler{
		Logger:        logger,
		Handler:       handler,
		Middleware:    middleware,
		xFrameOptions: xFrameOptions,
	}
}

type apiHandler struct {
	lager.Logger
	http.Handler
	token.Middleware

	xFrameOptions string
}

func (p *apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	tokenString := p.Middleware.GetAuthToken(r)
	if tokenString != "" {
		p.proxyWebRequest(w, r)
	}

	p.Handler.ServeHTTP(w, r)
}

func (p *apiHandler) proxyWebRequest(w http.ResponseWriter, r *http.Request) {

	tokenString := p.Middleware.GetAuthToken(r)
	if tokenString != "" {
		r.Header.Set("Authorization", tokenString)
	}

	if p.shouldCheckCSRF(r) {
		csrfHeader := r.Header.Get(CSRFHeaderName)
		if csrfHeader == "" {
			p.Logger.Debug("csrf-header-is-not-set")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		csrfToken := p.Middleware.GetCSRFToken(r)
		if csrfToken == "" {
			p.Logger.Debug("csrf-cookie-is-not-set")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if csrfToken != csrfHeader {
			p.Logger.Debug("csrf-token-does-not-match-auth-token", lager.Data{
				"auth-csrf-token":    csrfToken,
				"request-csrf-token": csrfHeader,
			})
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if p.xFrameOptions != "" {
		w.Header().Set("X-Frame-Options", p.xFrameOptions)
	}

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Download-Options", "noopen")
}

func (p *apiHandler) shouldCheckCSRF(r *http.Request) bool {
	switch r.Method {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		return false
	default:
		return true
	}
}
