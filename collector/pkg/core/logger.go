package core

import (
	"log"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggerConfig is the possible logger configurations
type LoggerConfig struct {
	Level string `envconfig:"LOG_LEVEL" default:"info"`
	Type  string `envconfig:"LOG_TYPE" default:"meteo-collector"`
}

// NewLogger builds and configures logger instance
func NewLogger(cfg LoggerConfig) (*zap.Logger, bool, error) {
	var err error
	logConfig := zap.NewProductionConfig()

	logLevel := new(zap.AtomicLevel)
	if err := logLevel.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil, false, err
	}

	isDebug := logLevel.String() == zapcore.DebugLevel.String()

	logConfig.Level = *logLevel
	logConfig.Development = isDebug
	logConfig.Sampling = nil
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logConfig.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	logConfig.InitialFields = map[string]interface{}{"type": cfg.Type}

	logger, err := logConfig.Build()
	if err != nil {
		return nil, false, err
	}

	// override std logger to write into app logger
	log.SetOutput(&stdLoggerWriter{logger, isDebug})

	return logger, isDebug, nil
}

type stdLoggerWriter struct {
	logger *zap.Logger
	debug  bool
}

func (w *stdLoggerWriter) Write(p []byte) (int, error) {
	msg := string(p)
	if w.debug && strings.Contains(msg, " Error when uploading spans to Jaeger") {
		// small hack to avoid Jeager error messages in debug mode like
		// 2019/11/25 16:19:44 Error when uploading spans to Jaeger: write udp 127.0.0.1:58948->127.0.0.1:6831: write: connection refused\n
		return len(p), nil
	}
	w.logger.Info(msg, zap.String("source", "std-log"))
	return len(p), nil
}
