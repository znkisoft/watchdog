package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/znkisoft/watchdog/pkg/ssm"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Server struct {
	ShutdownCtx context.Context
	BindAddress string
	Monitor     *ssm.Monitor
}

const (
	HeaderContentType = "Content-Type"

	MimeApplicationXProtobuf = "application/x-protobuf"
	MimeApplicationJSON      = "application/json"
)

func (s *Server) Start(ctx context.Context) error {
	routers := NewRouter([]mux.MiddlewareFunc{
		RequestLogMiddleware,
		ContextMiddleware,
	})

	s.BindAddress = ":3020"
	s.ShutdownCtx = ctx
	s.Monitor = ssm.RegisterMonitor()

	log.Printf("[INFO] [http,server] [message: starting HTTP server on port %s]", s.BindAddress)
	httpServer := &http.Server{
		Addr:    s.BindAddress,
		Handler: routers,
	}
	go shutdown(s.ShutdownCtx, httpServer)
	err := httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("[ERROR] [message: http server failed] [error: %s]", err)
	}

	return nil
}

func UnmarshalReq(r *http.Request, req proto.Message) error {
	ct := r.Header.Get(HeaderContentType)
	switch ct {
	case MimeApplicationJSON:
		return unmarshalJSON(r, req)
	case MimeApplicationXProtobuf:
		return unmarshalProto(r, req)
	}
	return unmarshalProto(r, req)
}

func unmarshalJSON(r *http.Request, req proto.Message) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	// log.Infof(r.Context(), "request body: %v", string(body))
	return protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(body, req)
}

func unmarshalProto(r *http.Request, req proto.Message) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := proto.Unmarshal(body, req); err != nil {
		return err
	}

	// ctx := r.Context()
	// jsonBytes, err := protojson.Marshal(req)
	// if err != nil {
	// 	log.Warningf(ctx, "failed to marshal proto message to json, err: %v", err)
	// } else {
	// 	log.Infof(ctx, "request body: %v", string(jsonBytes))
	// }

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
