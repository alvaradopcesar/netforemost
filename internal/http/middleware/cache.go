package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"netforemost/pkg/cache"
)

// WithCache response with data in cache if exists.
func WithCache(exp time.Duration, ca cache.Cache) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		if ca == nil {
			return next
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Cache-Control") == "no-cache" {
				next.ServeHTTP(w, r)
				return
			}

			switch r.Method {
			case http.MethodGet:
				readHandler(exp, ca, next, w, r)
			case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
				writeHandler(ca, next, w, r)
			default:
				next.ServeHTTP(w, r)
			}
		})
	}
}

func readHandler(exp time.Duration, ca cache.Cache, next http.Handler, w http.ResponseWriter, r *http.Request) {
	if b, err := ca.Get(r.URL.String()); err == nil && b != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(b)
		return
	}

	c := httptest.NewRecorder()
	next.ServeHTTP(c, r)
	for k, v := range c.Header() {
		w.Header()[k] = v
	}

	w.WriteHeader(c.Code)
	b := c.Body.Bytes()
	_, _ = w.Write(b)

	if c.Code != http.StatusOK {
		return
	}
	_ = ca.Set(r.URL.String(), b, exp)
}

func writeHandler(ca cache.Cache, next http.Handler, w http.ResponseWriter, r *http.Request) {
	next.ServeHTTP(w, r)
	sPath := strings.Split(r.URL.Path, "/")
	if len(sPath) < 2 {
		return
	}

	var pattern strings.Builder
	pattern.WriteString("*")
	pattern.WriteString(sPath[1])
	pattern.WriteString("*")

	keys, err := ca.Keys(pattern.String())
	if err != nil || len(keys) < 1 {
		return
	}

	err = ca.Del(keys...)
	if err != nil {
		return
	}
}
