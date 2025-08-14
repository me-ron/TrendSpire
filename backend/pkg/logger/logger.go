package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var log zerolog.Logger

func InitLogger(env string) {
	zerolog.TimeFieldFormat = time.RFC3339

	if env == "production" {
		log = zerolog.New(os.Stdout).With().Timestamp().Logger()
	} else {
		log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
			With().
			Timestamp().
			Logger()
	}
}

func Info(msg string, fields ...interface{}) {
	log.Info().Fields(toMap(fields)).Msg(msg)
}

func Error(msg string, fields ...interface{}) {
	log.Error().Fields(toMap(fields)).Msg(msg)
}

func Fatal(msg string, fields ...interface{}) {
	log.Fatal().Fields(toMap(fields)).Msg(msg)
}

func toMap(fields []interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for i := 0; i < len(fields)-1; i += 2 {
		key, ok := fields[i].(string)
		if ok {
			m[key] = fields[i+1]
		}
	}
	return m
}
