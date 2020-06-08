package logger

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	Info = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)

}

func SetupLogger(consoleWriter bool) {
	// Prod
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.TimestampFieldName = "created"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorFieldName = "message"
	zerolog.ErrorStackMarshaler = MarshalStack
	log.Logger = log.With().Caller().Logger()

	if consoleWriter {
		// Dev
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(ConsoleWriter{
			Out:           os.Stderr,
			NoColor:       false,
			TimeFormat:    "2006-01-02 15:04:05",
			MarshalIndent: true,
		})
	}
}
