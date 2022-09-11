package server

import (
	"net/http"
)

const (
	CtxKeyUser = "user"
)

func GetUserContext(r *http.Request) string {
	return r.Context().Value(CtxKeyUser).(string)
}
