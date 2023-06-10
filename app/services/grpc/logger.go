package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
)

func InterceptorLogger(l *zap.SugaredLogger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make([]interface{}, 0, len(fields))

		for i := 0; i < len(fields); i += 2 {
			key := fields[i]
			value := fields[i+1]
			f = append(f, key.(string), value.(string))
		}

		lgr := l.WithOptions(zap.AddCallerSkip(1)).With(f...)

		switch lvl {
		case logging.LevelDebug:
			lgr.Debug(msg)
		case logging.LevelInfo:
			lgr.Info(msg)
		case logging.LevelWarn:
			lgr.Warn(msg)
		case logging.LevelError:
			lgr.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}
