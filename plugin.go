// Package plugin adds X-Http-Version headers to incoming requests.
package plugin

import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct{}

// Demo a Demo plugin.
type Demo struct {
	next http.Handler
	name string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Demo{
		next: next,
		name: name,
	}, nil
}

func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.Header.Add("X-Http-Version", req.Proto)

	a.next.ServeHTTP(rw, req)
}
