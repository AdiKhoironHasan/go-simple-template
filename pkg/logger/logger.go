package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type logger struct {
	Logger zerolog.Logger
	layers []string
}

func NewLogger() *logger {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// Initialize logger with console output
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Set custom caller hook to include line number
	log.Logger = log.With().Caller().Logger()

	return &logger{
		Logger: log.Logger,
	}
}
