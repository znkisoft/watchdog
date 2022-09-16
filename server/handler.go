package server

import (
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type ProtoHandlerFunc func(http.ResponseWriter, *http.Request) (proto.Message, error)
type JSONHandlerFunc func(http.ResponseWriter, *http.Request) (any, error) // NOTE any?

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
			http.Error(w, "[internal error]: "+s.Err().Error(), http.StatusInternalServerError)
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

	// TODO add different content type to marshal

	// w.Header().Set(HeaderContentType, MimeApplicationXProtobuf)
	w.Header().Set(HeaderContentType, MimeApplicationJSON)

	jsonBytes, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(m)
	// protoBytes, err := proto.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(jsonBytes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (fn JSONHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = r.Context()
	m, err := fn(w, r)
	if err != nil {
		if se, ok := err.(*Err); ok {
			http.Error(w, se.Message, se.Code)
			return
		}

		s, ok := status.FromError(err)
		if !ok {
			http.Error(w, "[internal error]: "+s.Err().Error(), http.StatusInternalServerError)
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

	w.Header().Set(HeaderContentType, MimeApplicationJSON)
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(jsonBytes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
