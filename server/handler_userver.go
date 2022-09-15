package server

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/znkisoft/watchdog/pkg/db"
	"github.com/znkisoft/watchdog/schema"
	"google.golang.org/protobuf/proto"
)

func AddUserver(_ http.ResponseWriter, r *http.Request) (proto.Message, error) {
	req := &schema.Userver{}
	if err := UnmarshalReq(r, req); err != nil {
		return nil, err
	}
	m := db.NewUserModel(req)
	if err := m.Save(); err != nil {
		return nil, err
	}
	return nil, nil
}

func RetrieveUservers(_ http.ResponseWriter, r *http.Request) (proto.Message, error) {
	// TODO better way to get queries
	skip, _ := strconv.Atoi(r.URL.Query().Get("skip"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if skip < 0 || limit < 1 {
		return nil, BadRequest("skip and limit params are invalid")
	}

	req := &schema.Pagination{
		Skip:  int32(skip),
		Limit: int32(limit),
	}

	m := db.NewUserModel(&schema.Userver{})
	result, err := m.Get(req)
	if err != nil {
		return nil, err
	}

	return &schema.UserverResponse{
		Uservers: result,
	}, nil
}

func RetrieveUserver(_ http.ResponseWriter, r *http.Request) (proto.Message, error) {
	req := &schema.Userver{}
	if err := UnmarshalReq(r, req); err != nil {
		return nil, err
	}
	return nil, nil
}

func DeleteUserver(_ http.ResponseWriter, r *http.Request) (proto.Message, error) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, BadRequest("id is required")
	}

	m := db.NewUserModel(&schema.Userver{})
	if err := m.Delete(id); err != nil {
		return nil, err
	}

	return nil, nil
}
