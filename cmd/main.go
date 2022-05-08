package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/znkisoft/watchdog/server"
)

func main() {
	host := ""
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	routers := server.NewRouter([]mux.MiddlewareFunc{
		server.LogMiddleware,
	})
	srv := &http.Server{
		Handler: routers,
		Addr:    host + ":" + port,
	}

	log.Printf("server started at %s:%s", host, port)
	log.Fatal(srv.ListenAndServe())
}
