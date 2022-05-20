package server

import (
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	watchdog "github.com/znkisoft/watchdog/schema"
	"google.golang.org/protobuf/proto"
)

type httpMethod string

type ProtoRouterMap map[string]map[httpMethod]ProtoHandlerFunc
type JSONRouterMap map[string]map[httpMethod]JSONHandlerFunc

var protoRouterMap = ProtoRouterMap{
	"/hello": {
		http.MethodGet: func(w http.ResponseWriter, r *http.Request) (proto.Message, error) {
			p := &watchdog.Ping{Message: "pong"}
			return p, nil
		},
	},
	"/uptime": {http.MethodPost: Uptime},
}

var jsonRouterMap = JSONRouterMap{
	"/docker/containers":              {http.MethodGet: ContainerList},
	"/docker/containers/{id}/create":  {http.MethodPost: ContainerCreate},
	"/docker/containers/{id}/stop":    {http.MethodPost: ContainerStop},
	"/docker/containers/{id}/restart": {http.MethodPost: ContainerRestart},
	"/docker/containers/{id}/kill":    {http.MethodPost: ContainerKill},
	"/docker/containers/{id}/remove":  {http.MethodDelete: ContainerRemove},
}

func NewRouter(mws []mux.MiddlewareFunc) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/").Subrouter()

	api.Methods(http.MethodOptions).Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	api.Use(mws...)

	// AMEND: better way to handle different types of routes
	protoRouters := make([]string, 0, len(protoRouterMap))
	for k := range protoRouterMap {
		protoRouters = append(protoRouters, k)
	}
	sort.Strings(protoRouters)
	for _, pt := range protoRouters {
		hs := protoRouterMap[pt]
		for m, hm := range hs {
			api.Path(pt).Methods(string(m)).Handler(hm)
		}
	}

	jsonRouters := make([]string, 0, len(jsonRouterMap))
	for k := range jsonRouterMap {
		jsonRouters = append(jsonRouters, k)
	}
	sort.Strings(jsonRouters)
	for _, pt := range jsonRouters {
		hs := jsonRouterMap[pt]
		for m, hm := range hs {
			api.Path(pt).Methods(string(m)).Handler(hm)
		}
	}
	return r
}
