package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/vgarvardt/meteo/collector/pkg/consume"
	"github.com/vgarvardt/meteo/collector/pkg/core"
)

// NewConsumeCmd creates a new consume command
func NewConsumeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "consume",
		Short: "Starts the consumer that listens to MQTTConfig broken messages",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			var cfg consume.Config
			if err := envconfig.Process("", &cfg); err != nil {
				return fmt.Errorf("could not load config from env: %w", err)
			}

			logger, _, err := core.NewLogger(cfg.LoggerConfig)
			if err != nil {
				return fmt.Errorf("could not build logger instance: %w", err)
			}

			mqttClient, err := consume.NewMQTTClient(cfg.MQTTConfig, logger)
			if err != nil {
				return fmt.Errorf("could not connect to MQTT broker: %w", err)
			}
			defer mqttClient.Disconnect(uint(cfg.DisconnectTimeout / time.Millisecond))

			if err := consume.StartConsumer(mqttClient, cfg.MQTTConfig.TopicSensors, consume.NewHandlerClimate(ctx, logger)); err != nil {
				logger.Error("Sensors consumer subscription finished with error", zap.Error(err))
				return err
			}

			if err := consume.StartConsumer(mqttClient, cfg.MQTTConfig.TopicSystem, consume.NewHandlerSystem(ctx, logger)); err != nil {
				logger.Error("System consumer subscription finished with error", zap.Error(err))
				return err
			}

			osSignal := make(chan os.Signal)
			signal.Notify(osSignal, os.Interrupt)

			select {
			case sig := <-osSignal:
				logger.Info("Got OS signal, stopping", zap.String("signal", sig.String()))
			}

			return nil
		},
	}
}
