package server

import (
	"github.com/binbinly/pkg/transport/http"
	"project-layout/internal/app"
)

// NewHTTPServer http server
func NewHTTPServer(c *app.ServerConfig) *http.Server {
	srv := http.NewServer(
		http.WithAddress(c.Addr),
		http.WithReadTimeout(c.ReadTimeout),
		http.WithWriteTimeout(c.WriteTimeout),
	)

	return srv
}
