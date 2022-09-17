package ssm

import (
	"github.com/znkisoft/watchdog/schema"
)

type Monitor struct {
	Config  string
	Plugins map[string]*Plugin
	// TODO for consistently monitor
}

type UnitMap map[string]string

var unitMap = UnitMap{
	"percent":  "float",
	"percents": "float",
	"number":   "int",
	"numbers":  "int",
	"int":      "int",
	"ints":     "int",
	"float":    "float",
	"floats":   "float",
	"second":   "int",
	"seconds":  "int",
	"byte":     "int",
	"bytes":    "int",
}

// Plugin is the general struct for all plugins
// it included the name, unit, config, status, output
type Plugin struct {
	Type   schema.PluginType
	Unit   string
	Enable bool
}

func RegisterMonitor() *Monitor {
	m := make(map[string]*Plugin, len(schema.PluginType_name))

	m[schema.PluginType_CPU.String()] = registerPlugin(schema.PluginType_CPU, "")
	m[schema.PluginType_UPTIME.String()] = registerPlugin(schema.PluginType_UPTIME, "")

	return &Monitor{
		Plugins: m,
	}
}

func registerPlugin(typ schema.PluginType, unit string) *Plugin {
	return &Plugin{
		Type: typ,
		Unit: "",
	}
}
