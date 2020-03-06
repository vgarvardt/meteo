package cmd

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/kelseyhightower/envconfig"
	wErrors "github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/vgarvardt/meteo/collector/pkg/consume"
	"github.com/vgarvardt/meteo/collector/pkg/core"
)

// NewConsumeCmd creates a new consume command
func NewConsumeCmd(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "consume",
		Short: "Starts the consumer that listens to MQTTConfig broken messages",
		RunE: func(cmd *cobra.Command, args []string) error {
			var cfg consume.Config
			if err := envconfig.Process("", &cfg); err != nil {
				return wErrors.Wrap(err, "could not load config from env")
			}

			logger, _, err := core.NewLogger(cfg.LoggerConfig)
			if err != nil {
				return wErrors.Wrap(err, "could not build logger instance")
			}

			mqttClient, err := consume.NewClient(cfg.MQTTConfig, logger)
			if err != nil {
				return wErrors.Wrap(err, "could not connect to MQTT broker")
			}
			defer mqttClient.Disconnect(uint(cfg.DisconnectTimeout / time.Millisecond))

			ctxWithCancel, cancel := context.WithCancel(ctx)
			defer cancel()

			var consumerErr error
			go func() {
				consumerErr = consume.RunConsumer(ctxWithCancel, mqttClient, cfg.MQTTConfig.Topics, logger)
			}()

			osSignal := make(chan os.Signal)
			signal.Notify(osSignal, os.Interrupt)

			select {
			case sig := <-osSignal:
				logger.Info("Got OS signal, stopping", zap.String("signal", sig.String()))
			}

			return consumerErr
		},
	}
}
