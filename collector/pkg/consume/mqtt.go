package consume

import (
	"fmt"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MQTTConfig struct {
	Server   string `envconfig:"MQTT_HOST" default:"tcp://127.0.0.1:1883"`
	ClientID string `envconfig:"MQTT_CLIENT_ID" default:"meteo-collector"`
	Topics   string `envconfig:"MQTT_TOPICS" default:"home/#"`

	ConnectTimeout    time.Duration `envconfig:"MQTT_CONNECT_TIMEOUT" default:"3s"`
	DisconnectTimeout time.Duration `envconfig:"MQTT_DISCONNECT_TIMEOUT" default:"3s"`
}

func NewClient(cfg MQTTConfig, logger *zap.Logger) (mqtt.Client, error) {
	mqtt.ERROR = mqttLogger{
		logger: logger,
		level:  zapcore.ErrorLevel,
	}
	mqtt.CRITICAL = mqttLogger{
		logger: logger,
		level:  zapcore.ErrorLevel,
	}
	mqtt.WARN = mqttLogger{
		logger: logger,
		level:  zapcore.WarnLevel,
	}
	mqtt.DEBUG = mqttLogger{
		logger: logger,
		level:  zapcore.DebugLevel,
	}

	client := mqtt.NewClient(mqtt.NewClientOptions().
		SetAutoReconnect(true).
		SetCleanSession(false).
		SetClientID(cfg.ClientID).
		SetOnConnectHandler(onConnectHandler(logger)).
		SetConnectionLostHandler(connectionLostHandler(logger)).
		AddBroker(cfg.Server))

	token := client.Connect()
	for !token.WaitTimeout(cfg.ConnectTimeout) {
	}

	return client, token.Error()
}

func onConnectHandler(logger *zap.Logger) mqtt.OnConnectHandler {
	return func(c mqtt.Client) {
		logger.Info("Connected to MQTT Broker")
	}
}

func connectionLostHandler(logger *zap.Logger) mqtt.ConnectionLostHandler {
	return func(client mqtt.Client, err error) {
		logger.Warn("Lost connection to MQTT Broker", zap.Error(err))
	}
}

type mqttLogger struct {
	logger *zap.Logger
	level  zapcore.Level
}

func (l mqttLogger) Println(v ...interface{}) {
	l.log(l.logger.With(zap.Any("data", v)), "MQTT client log message")
}

func (l mqttLogger) Printf(format string, v ...interface{}) {
	l.log(
		l.logger.With(
			zap.String("msg", fmt.Sprintf(format, v...)),
			zap.Any("data", v),
			zap.String("format", format),
		),
		"MQTT client log formatted message",
	)
}

func (l mqttLogger) log(ll *zap.Logger, msg string) {
	switch l.level {
	case zapcore.ErrorLevel:
		ll.Error(msg)
	case zapcore.DebugLevel:
		ll.Debug(msg)
	case zapcore.WarnLevel:
		fallthrough
	default:
		ll.Warn(msg)
	}
}
