package server

import (
	"net/http"

	"github.com/znkisoft/watchdog/schema"
	"google.golang.org/protobuf/proto"
)

func Uptime(_ http.ResponseWriter, r *http.Request) (proto.Message, error) {
	req := &schema.PluginRequest{}
	if err := UnmarshalReq(r, req); err != nil {
		return nil, err
	}

	if req.Type != schema.PluginType_UPTIME {
		return nil, nil
	}

	return &schema.PluginResponse{
		PluginInfo: &schema.PluginResponse_UptimeInfo{
			UptimeInfo: &schema.UptimeInfo{
				Seconds: 12313,
			},
		},
	}, nil
}
