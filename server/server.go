package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	ShutdownCtx context.Context
	BindAddress string
}

func (server *Server) Start(ctx context.Context) error {
	// initialize the routes and middleware
	routers := NewRouter([]mux.MiddlewareFunc{
		LogMiddleware,
	})

	server.BindAddress = ":8080"
	server.ShutdownCtx = ctx

	log.Printf("[INFO] [http,server] [message: starting HTTP server on port %s]", server.BindAddress)
	httpServer := &http.Server{
		Addr:    server.BindAddress,
		Handler: routers,
	}
	go shutdown(server.ShutdownCtx, httpServer)
	err := httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("[ERROR] [message: http server failed] [error: %s]", err)
	}

	return nil
}

func shutdown(shutdownCtx context.Context, httpServer *http.Server) {
	<-shutdownCtx.Done()

	log.Println("[DEBUG] [server,server] [message: shutting down server server]")
	shutdownTimeout, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := httpServer.Shutdown(shutdownTimeout)
	if err != nil {
		fmt.Printf("[ERROR] [server,server] [message: failed shutdown server server] [error: %s]", err)
	}
}
