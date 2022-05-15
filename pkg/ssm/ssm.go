package ssm

import (
	"github.com/znkisoft/watchdog/schema"
)

type Monitor struct {
	Config  string
	Plugins map[string]*Plugin
	// TODO for consistently monitor
}

// Plugin is the general struct for all plugins
// it included the name, unit, config, status, output
type Plugin struct {
	Type schema.PluginType
	Unit string // TODO define unit type and unit number
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
