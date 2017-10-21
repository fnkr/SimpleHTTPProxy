package main

import (
    "github.com/elazarl/goproxy"
    "log"
    "fmt"
    "net/http"
)

func main() {
    config := getConfig()
    proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = config.Verbose
    log.Print(fmt.Sprintf("Listening on %s:%d...", config.Host, config.Port))
    log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, config.Port), proxy))
}
