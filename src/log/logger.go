package log

import (
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

type Loger struct {
	Log zerolog.Logger
}

func init() {
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	logger = zerolog.New(output).With().Timestamp().Logger()
}

func Log(level zerolog.Level, message string) {
	logger.WithLevel(level).Msg(message)
}
