package proxyhandler

import (
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/skymarshal/token"
)

func NewAuthHandler(
	logger lager.Logger,
	target *url.URL,
	middleware token.Middleware,
) *authHandler {
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

	return &authHandler{
		Logger:     logger,
		Handler:    handler,
		Middleware: middleware,
	}
}

type authHandler struct {
	lager.Logger
	http.Handler
	token.Middleware
}

func (p *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Host == "" {
		p.proxyWebRequest(w, r)
	}

	p.Handler.ServeHTTP(w, r)
}

func (p *authHandler) proxyWebRequest(w http.ResponseWriter, r *http.Request) {

	tokenString := p.Middleware.GetAuthToken(r)
	if tokenString != "" {
		r.Header.Set("Authorization", tokenString)
	}
}
