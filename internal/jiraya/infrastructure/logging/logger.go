package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
	"strings"
)

var logger *zap.Logger

func Get() *zap.Logger {
	if logger == nil {
		err := os.MkdirAll("logs", 0755)

		if err != nil || os.IsExist(err) {
			panic("failed to create logs folder")
		} else {
			//file, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
			//
			//if err != nil {
			//	panic(fmt.Sprintf("cannot open logs file: %s", err))
			//}

			logger = New(os.Getenv("log_level"), os.Stdout)
		}
	}

	return logger
}

// New returns a new logger instance. If out is nil, os.Stdout will be used.
func New(logLevel string, out io.Writer) *zap.Logger {
	if out == nil {
		out = os.Stdout
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(out),
		getZapLogLevel(logLevel),
	)

	l := zap.New(core)

	return l
}

func getZapLogLevel(logLevel string) zapcore.Level {
	switch strings.ToLower(logLevel) {
	case `debug`:
		return zapcore.DebugLevel
	case `info`:
		return zapcore.InfoLevel
	case `warn`:
		return zapcore.WarnLevel
	case `error`:
		return zapcore.ErrorLevel
	case `panic`:
		return zapcore.PanicLevel
	default:
		log.Println("UNKNOWN LOGGER LOG LEVEL. USING DEBUG LEVEL AS DEFAULT.")
		return zapcore.DebugLevel
	}
}
