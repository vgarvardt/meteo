package consume

import "github.com/vgarvardt/meteo/collector/pkg/core"

type Config struct {
	MQTTConfig
	core.LoggerConfig
}
