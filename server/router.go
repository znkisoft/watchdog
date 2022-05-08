package server

import (
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	watchdog "github.com/znkisoft/watchdog/schema"
	"google.golang.org/protobuf/proto"
)

type httpMethod string

type RouterMap map[string]map[httpMethod]ProtoHandlerFunc

var routerMap = RouterMap{
	"/hello": {
		http.MethodGet: func(w http.ResponseWriter, r *http.Request) (proto.Message, error) {
			p := &watchdog.Ping{Message: "pong"}
			return p, nil
		},
	},
}

func NewRouter(mws []mux.MiddlewareFunc) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/").Subrouter()

	api.Methods(http.MethodOptions).Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	api.Use(mws...)

	routers := make([]string, 0, len(routerMap))
	for k := range routerMap {
		routers = append(routers, k)
	}
	sort.Strings(routers)
	for _, pt := range routers {
		hs := routerMap[pt]
		for m, hm := range hs {
			api.Path(pt).Methods(string(m)).Handler(hm)
		}
	}
	return r
}
