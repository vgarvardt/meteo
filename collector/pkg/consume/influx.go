package consume

import (
	"context"
	"io"

	"github.com/influxdata/influxdb-client-go"
	wErrors "github.com/pkg/errors"
	"go.uber.org/zap"
)

type InfluxDBConfig struct {
	Server string `envconfig:"INFLUX_SERVER" default:"https://eu-central-1-1.aws.cloud2.influxdata.com"`
	Token  string `envconfig:"INFLUX_TOKEN"`
	Org    string `envconfig:"INFLUX_ORG" required:"true"`
}

type InfluxDBClient interface {
	io.Closer
	Write(ctx context.Context, bucket, org string, m ...influxdb.Metric) (n int, err error)
}

func NewInfluxDBClient(ctx context.Context, cfg InfluxDBConfig, logger *zap.Logger) (InfluxDBClient, error) {
	if cfg.Token == "" {
		logger.Warn("InfluxDB token is not set, falling back to logger client")
		return &loggerInfluxClient{logger: logger}, nil
	}

	influx, err := influxdb.New(cfg.Server, cfg.Token)
	if err != nil {
		return nil, wErrors.Wrap(err, "could not instantiate InfluxDB client")
	}

	if err := influx.Ping(ctx); err != nil {
		return nil, wErrors.Wrap(err, "could not ping InfluxDB server")
	}

	return &orgAwareInfluxClient{
		org:    cfg.Org,
		client: influx,
	}, nil
}

type loggerInfluxClient struct {
	logger *zap.Logger
}

func (c *loggerInfluxClient) Close() error {
	return nil
}

// Write writes metrics to a bucket, and org.  It retries intelligently.
// If the write is too big, it retries again, after breaking the payloads into two requests.
func (c *loggerInfluxClient) Write(ctx context.Context, bucket, org string, m ...influxdb.Metric) (n int, err error) {
	c.logger.Info(
		"Sending Influx metrics to log instead of InfluxDB",
		zap.String("bucket", bucket), zap.String("org", org), zap.Any("m", m),
	)

	return len(m), nil
}

// orgAwareInfluxClient ignores org parameter in the Write method and uses internal one
type orgAwareInfluxClient struct {
	org    string
	client InfluxDBClient
}

func (c *orgAwareInfluxClient) Write(ctx context.Context, bucket, org string, m ...influxdb.Metric) (n int, err error) {
	return c.client.Write(ctx, bucket, c.org, m...)
}

func (c *orgAwareInfluxClient) Close() error {
	return c.client.Close()
}
