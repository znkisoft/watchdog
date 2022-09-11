package server

import (
	"net/http"

	"github.com/znkisoft/watchdog/schema"
	"google.golang.org/protobuf/proto"
)

func AddUserver(_ http.ResponseWriter, r *http.Request) (proto.Message, error) {
	req := &schema.Userver{}
	if err := UnmarshalReq(r, req); err != nil {
		return nil, err
	}
	return nil, nil
}

func RetrieveUservers(_ http.ResponseWriter, r *http.Request) (proto.Message, error) {
	req := &schema.Userver{}
	if err := UnmarshalReq(r, req); err != nil {
		return nil, err
	}

	return nil, nil
}

func RetrieveUserver(_ http.ResponseWriter, r *http.Request) (proto.Message, error) {
	req := &schema.Userver{}
	if err := UnmarshalReq(r, req); err != nil {
		return nil, err
	}
	return nil, nil
}

func DeleteUserver(_ http.ResponseWriter, r *http.Request) (proto.Message, error) {
	req := &schema.Userver{}
	if err := UnmarshalReq(r, req); err != nil {
		return nil, err
	}
	return nil, nil
}
