package server

import (
	"context"
	"log"
	"net/http"
)

func RequestLogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}

func ContextMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO replace it with real context
		token := "token"
		ctx := context.WithValue(r.Context(), CtxKeyUser, token)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
