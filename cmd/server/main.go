package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kazan/app-skeleton/app"
)

type Config struct {
	Server Server
}

type Server struct {
	Addr              string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

func ConfigDefaults() *Config {
	return &Config{
		Server: Server{
			Addr:              ":8080",
			ReadTimeout:       15 * time.Second,
			ReadHeaderTimeout: 15 * time.Second,
			WriteTimeout:      15 * time.Second,
			IdleTimeout:       15 * time.Second,
		},
	}
}

func main() {
	// read config
	cfg := ConfigDefaults()

	// gin global setup
	r := gin.New()
	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// setup routes
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": "1.0.0"})
	})

	// setup dependencies
	if err := app.Load(r); err != nil {
		panic(err)
	}

	// setup server
	server := server(cfg.Server, r)

	// start erver
	panic(server.ListenAndServe())
}

func server(cfg Server, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              cfg.Addr,
		Handler:           handler,
		ReadTimeout:       cfg.ReadTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}
}
