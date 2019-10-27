package logging

import (
	"encoding/json"

	zap "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger will return a standard logger with pre-customised logging configurations
func NewLogger() *zap.SugaredLogger {
	rawJSON := []byte(`{
		"level": "info",
		"development": true,
		"encoding": "console",
		"outputPaths": ["stdout", "/tmp/logs"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
			"messageKey": "m",
			"levelKey": "l",
			"timeKey": "t",
			"levelEncoder": "capitalColor"
		}
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	sugar := logger.Sugar()

	return sugar
}
