package consume

import (
	"context"
	"encoding/json"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	wErrors "github.com/pkg/errors"
	"go.uber.org/zap"
)

type msg struct {
	When time.Time `json:"when"`
	Data float32   `json:"data"`
}

func RunConsumer(ctx context.Context, mqttClient mqtt.Client, topicSensors string, logger *zap.Logger) error {
	token := mqttClient.Subscribe(topicSensors, 1, messageHandler(ctx, logger))
	for !token.Wait() {
		select {
		case <-ctx.Done():
			logger.Info("Got Done signal, quitting consumer")
			return nil
		}
	}

	if err := token.Error(); err != nil {
		return wErrors.Wrap(token.Error(), "something went wrong with MQTT broker subscription")
	}

	return token.Error()
}

func messageHandler(ctx context.Context, logger *zap.Logger) mqtt.MessageHandler {
	return func(client mqtt.Client, message mqtt.Message) {
		ll := logger.With(
			zap.String("topic", message.Topic()),
			zap.Int("qos", int(message.Qos())),
			zap.Binary("payload-b", message.Payload()),
			zap.ByteString("payload-s", message.Payload()),
		)

		ll.Debug("Consumer received a message")

		var m msg
		if err := json.Unmarshal(message.Payload(), &m); err != nil {
			ll.Error("Could not unmarshal incoming message", zap.Error(err))
			message.Ack()
			return
		}

		ll.Debug("Decoded incoming message", zap.Any("msg", m))

		message.Ack()
	}
}
