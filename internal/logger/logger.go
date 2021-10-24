package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func New() zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	return logger
}

func NewWithPrefix(key, value string) zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	return logger.With().Str(key, value).Logger()
}
