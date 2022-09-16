package server

import (
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	"github.com/znkisoft/watchdog/schema"
	"google.golang.org/protobuf/proto"
)

type httpMethod string

type ProtobufRouterMap map[string]map[httpMethod]ProtoHandlerFunc
type JSONRouterMap map[string]map[httpMethod]JSONHandlerFunc

var pbRouterMap = ProtobufRouterMap{
	"/health": {
		http.MethodGet: func(w http.ResponseWriter, r *http.Request) (proto.Message, error) {
			p := &schema.UptimeInfo{}
			return p, nil
		},
	},
	"/uptime": {http.MethodPost: Uptime},
	"/userver": {
		http.MethodGet:  RetrieveUservers,
		http.MethodPost: AddUserver,
	},
	"/userver/{id}": {
		http.MethodGet:    RetrieveUserver,
		http.MethodDelete: DeleteUserver,
	},
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
	api := r.PathPrefix("/api").Subrouter()

	api.Methods(http.MethodOptions).Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	api.Use(mws...)

	// AMEND: better way to handle different types of routes
	protoRouters := make([]string, 0, len(pbRouterMap))
	for k := range pbRouterMap {
		protoRouters = append(protoRouters, k)
	}
	sort.Strings(protoRouters)
	for _, pt := range protoRouters {
		hs := pbRouterMap[pt]
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

	api.Path("/ws").HandlerFunc(WebSocketHandler)
	return r
}
