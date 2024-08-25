// File: log/log.go

package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var log zerolog.Logger

func init() {
	log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).With().Timestamp().Caller().Logger()

}

// fields is a map of field names to values
type Fields map[string]interface{}

// Info returns a new Info event logger
func Info() *zerolog.Event {
	return log.Info()
}

// Error returns a new Error event logger
func Error() *zerolog.Event {
	return log.Error()
}

// Debug returns a new Debug event logger
func Debug() *zerolog.Event {
	return log.Debug()
}

// Warn returns a new Warn event logger
func Warn() *zerolog.Event {
	return log.Warn()
}

// Fatal returns a new Fatal event logger
func Fatal() *zerolog.Event {
	return log.Fatal()
}
