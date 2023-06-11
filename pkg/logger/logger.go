package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type logger struct {
	Logger zerolog.Logger
}

func NewLogger() *logger {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// Initialize logger with console output
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	})

	// Set custom caller hook to include line number
	log.Logger = log.With().Caller().Logger()

	return &logger{
		Logger: log.Logger,
	}
}
