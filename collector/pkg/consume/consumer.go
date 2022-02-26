package consume

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type msg struct {
	When          time.Time              `json:"when"`
	Data          map[string]interface{} `json:"data"`
	MeasurementID string                 `json:"measurement_id"`
}

func StartConsumer(mqttClient mqtt.Client, topic string, messageHandler mqtt.MessageHandler) error {
	token := mqttClient.Subscribe(topic, 1, messageHandler)
	for !token.Wait() {
	}

	if err := token.Error(); err != nil {
		return fmt.Errorf("something went wrong with MQTT broker subscription: %w", err)
	}

	return nil
}
