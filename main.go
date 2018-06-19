package main

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"errors"
)

var config *Options

func checkRequest(ctx *goproxy.ProxyCtx) error {
	auth := parseProxyAuthorization(ctx.Req.Header.Get("Proxy-Authorization"))
	if !checkUser(auth) {
		if config.Verbose {
			ctx.Logf("Request denied because authorization failed: \"%s\" => \"%s:%s\"", ctx.Req.Header.Get("Proxy-Authorization"), auth.Username, auth.Password)
		}
		return errors.New("invalid user")
	}

	if !checkHost(ctx.Req.Host) {
		if config.Verbose {
			ctx.Logf("Request denied because host is not allowed: %s", ctx.Req.Host)
		}
		return errors.New("invalid host")
	}

	return nil
}

func main() {
	config = parseArgs()

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = config.Verbose

	// http:// (GET, ...)
	proxy.OnRequest().DoFunc(func (req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		if err := checkRequest(ctx); err != nil {
			return req, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusForbidden, "403 Forbidden\n")
		}

		return req, nil
	})

	// https:// (CONNECT)
	proxy.OnRequest().HandleConnectFunc(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		if err := checkRequest(ctx); err != nil {
			return goproxy.RejectConnect, host
		}

		return goproxy.OkConnect, host
	})

	log.Fatal(http.ListenAndServe(config.Bind, proxy))
}
