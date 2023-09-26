package log

import (
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LogFormatter struct {
	LogLevel       string
	LevelFieldName string

	JSON            bool
	TimeFieldFormat string
	TimeFieldName   string
}

func NewLogger(serviceName string, opts ...LogFormatter) zerolog.Logger {
	logger := log.With().Str("service", strings.ToLower(serviceName))
	return logger.Logger()
}
