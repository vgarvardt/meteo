package consume

import (
	"context"

	"github.com/eclipse/paho.mqtt.golang"
	wErrors "github.com/pkg/errors"
	"go.uber.org/zap"
)

func RunConsumer(ctx context.Context, mqttClient mqtt.Client, topic string, logger *zap.Logger) error {
	token := mqttClient.Subscribe(topic, 1, messageHandler(ctx, logger))
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
		logger.Debug(
			"Consumer received a message",
			zap.String("topic", message.Topic()),
			zap.Int("qos", int(message.Qos())),
			zap.Binary("payload-b", message.Payload()),
			zap.ByteString("payload-s", message.Payload()),
		)

		message.Ack()
	}
}
