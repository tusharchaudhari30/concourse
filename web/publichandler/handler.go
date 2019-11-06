package publichandler

import (
	"net/http"

	"github.com/gobuffalo/packr"
)

func NewHandler() http.Handler {
	return CacheNearlyForever(http.StripPrefix("/public/", http.FileServer(packr.NewBox("../public"))))
}
