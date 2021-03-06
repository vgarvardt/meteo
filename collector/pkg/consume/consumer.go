package consume

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/influxdata/influxdb-client-go"
	wErrors "github.com/pkg/errors"
	"go.uber.org/zap"
)

type msg struct {
	When          time.Time              `json:"when"`
	Data          map[string]interface{} `json:"data"`
	MeasurementID string                 `json:"measurement_id"`
}

func StartConsumer(ctx context.Context, mqttClient mqtt.Client, influx InfluxDBClient, topic string, logger *zap.Logger) error {
	token := mqttClient.Subscribe(topic, 1, messageHandler(ctx, influx, logger))
	for !token.Wait() {
	}

	if err := token.Error(); err != nil {
		return wErrors.Wrap(token.Error(), "something went wrong with MQTT broker subscription")
	}

	return nil
}

func messageHandler(ctx context.Context, influx InfluxDBClient, logger *zap.Logger) mqtt.MessageHandler {
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

		topicParts := strings.Split(message.Topic(), "/")
		if len(topicParts) != 3 {
			ll.Error("Invalid message topic format, expected <bucket>/<room>/<metricName>", zap.String("topic", message.Topic()))
			message.Ack()
			return
		}

		bucket, room, metricName := topicParts[0], topicParts[1], topicParts[2]

		metrics := []influxdb.Metric{
			influxdb.NewRowMetric(
				m.Data,
				metricName,
				map[string]string{"room": room},
				m.When,
			),
		}

		if _, err := influx.Write(ctx, bucket, "acme", metrics...); err != nil {
			ll.Error("Could not write metrics to influx", zap.Error(err), zap.Any("metrics", metrics))
		}

		message.Ack()
	}
}
