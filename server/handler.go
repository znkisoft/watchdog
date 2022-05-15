package server

import (
	"log"
	"net/http"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type ProtoHandlerFunc func(http.ResponseWriter, *http.Request) (proto.Message, error)

func (fn ProtoHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = r.Context()
	m, err := fn(w, r)
	if err != nil {
		if se, ok := err.(*Err); ok {
			http.Error(w, se.Message, se.Code)
			return
		}

		s, ok := status.FromError(err)
		if !ok {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		statusCode := HTTPStatusFromCode(s.Code())

		m = s.Proto()
		w.WriteHeader(statusCode)
	}

	if m == nil {
		return
	}

	// debug only
	log.Printf("[DEBUG] response from %s: %v", r.URL.String(), m)

	w.Header().Set(HeaderContentType, MimeApplicationXProtobuf)
	protoBytes, err := proto.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(protoBytes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
