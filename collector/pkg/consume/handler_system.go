package consume

import (
	"context"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	v1 "github.com/vgarvardt/meteo/collector/pkg/pb/measurement/v1"
)

func NewHandlerSystem(_ context.Context, logger *zap.Logger) mqtt.MessageHandler {
	return func(client mqtt.Client, message mqtt.Message) {
		// acknowledge message in any case
		defer message.Ack()

		ll := logger.With(
			zap.String("topic", message.Topic()),
			zap.Int("qos", int(message.Qos())),
		)

		ll.Debug("Consumer received a message")

		m := new(v1.System)
		if err := proto.Unmarshal(message.Payload(), m); err != nil {
			ll.Error("Could not unmarshal payload", zap.Error(err))
			return
		}

		ll.Debug("Decoded incoming message", zap.Any("msg", m))
	}
}
