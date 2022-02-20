package consume

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

type msg struct {
	When          time.Time              `json:"when"`
	Data          map[string]interface{} `json:"data"`
	MeasurementID string                 `json:"measurement_id"`
}

func StartConsumer(ctx context.Context, mqttClient mqtt.Client, topic string, logger *zap.Logger) error {
	token := mqttClient.Subscribe(topic, 1, messageHandler(ctx, logger))
	for !token.Wait() {
	}

	if err := token.Error(); err != nil {
		return fmt.Errorf("something went wrong with MQTT broker subscription: %w", err)
	}

	return nil
}

func messageHandler(_ context.Context, logger *zap.Logger) mqtt.MessageHandler {
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
