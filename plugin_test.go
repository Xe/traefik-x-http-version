package plugin_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	plugin "git.xeserv.us/xe/traefik-x-http-version"
)

func TestDemo(t *testing.T) {
	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := plugin.New(ctx, next, nil, "demo-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Proto = "HTTP/2.0"

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "X-Http-Version", "HTTP/2.0")
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}
