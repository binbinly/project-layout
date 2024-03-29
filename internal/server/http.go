package server

import (
	"github.com/binbinly/gin-pkg/app"
	"github.com/binbinly/pkg/transport/http"
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
