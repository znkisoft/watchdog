package server

import (
	"net/http"

	"github.com/znkisoft/watchdog/pkg/ssm"
	"github.com/znkisoft/watchdog/schema"
	"google.golang.org/protobuf/proto"
)

func ssh(_ http.ResponseWriter, r *http.Request) (proto.Message, error) {
	req := &schema.PluginRequest{}
	if err := UnmarshalReq(r, req); err != nil {
		return nil, err
	}

	data, err := ssm.Uptime()
	if err != nil {
		return nil, err
	}

	return &schema.PluginResponse{
		PluginInfo: &schema.PluginResponse_UptimeInfo{
			UptimeInfo: &schema.UptimeInfo{
				Seconds: int32(data),
			},
		},
	}, nil
}
