package internal

import "github.com/ooesili/dead-drop/internal/config"

var ConfigSchema = config.Schema{
	"PORT":      config.Optional("3000"),
	"BIND_ADDR": config.Optional("0.0.0.0"),
}
