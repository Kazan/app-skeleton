package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kazan/app-skeleton/cmd/server/config"
)

type module interface {
}

func main() {
	// read config
	cfg := config.NewWithDefaults()

	// setup dependencies

	// setup gin
	r := gin.New()
	r.Use(gin.Recovery())

	// setup routes

	// setup server
	server := server(cfg.Server, http.DefaultServeMux)

	// start erver
	panic(server.ListenAndServe())
}

func server(cfg config.Server, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              cfg.Addr,
		Handler:           handler,
		ReadTimeout:       cfg.ReadTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}
}
