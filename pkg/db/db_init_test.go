package db

import (
	"testing"

	"github.com/znkisoft/watchdog/schema"
	"google.golang.org/protobuf/proto"
)

func TestCreateDB(t *testing.T) {
	type args struct {
		schemas []proto.Message
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"return empty string with empty schemas",
			args{[]proto.Message{}},
			"",
		},
		{
			"return empty string with nil schemas",
			args{nil},
			"",
		}, {
			"get expected DDL",
			args{
				[]proto.Message{
					&schema.Userver{
						Id:            "",
						Name:          "",
						Ip:            "",
						Port:          "",
						Protocol:      "",
						CheckInterval: 0,
						Timeout:       0,
					}, &schema.Userver{
						Id:            "",
						Name:          "",
						Ip:            "",
						Port:          "",
						Protocol:      "",
						CheckInterval: 0,
						Timeout:       0,
					},
				},
			},
			"CREATE TABLE IF NOT EXISTS userver ( id TEXT PRIMARY KEY, name TEXT, ip TEXT, port TEXT, protocol TEXT, check_interval INTEGER, timeout INTEGER);\nCREATE TABLE IF NOT EXISTS userver ( id TEXT PRIMARY KEY, name TEXT, ip TEXT, port TEXT, protocol TEXT, check_interval INTEGER, timeout INTEGER);\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateDBQuery(tt.args.schemas); got != tt.want {
				t.Errorf("CreateDBQuery() and wanted:\n%s\n%s", got, tt.want)
			}
		})
	}
}
