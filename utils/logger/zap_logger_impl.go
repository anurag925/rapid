package logger

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapLogger struct {
	logger *zap.Logger
}

var _ Logger = (*ZapLogger)(nil)

func NewZapLogger(env string) *ZapLogger {
	var logger *zap.Logger
	// Set up logging configuration based on environment
	if env == "development" {
		// Development environment
		encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		consoleDebugging := zapcore.Lock(zapcore.AddSync(os.Stdout))

		fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		fileDebugging := zapcore.AddSync(&lumberjack.Logger{
			Filename:   "logs/application.log",
			MaxSize:    100, // MB
			MaxBackups: 3,
			MaxAge:     28, // Days
			LocalTime:  true,
		})

		consoleCore := zapcore.NewCore(encoder, consoleDebugging, zapcore.DebugLevel)
		fileCore := zapcore.NewCore(fileEncoder, fileDebugging, zapcore.DebugLevel)

		// Use a TeeEncoder to write to both the console and file
		logger = zap.New(zapcore.NewTee(consoleCore, fileCore), zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	} else {
		// Production or other environment
		encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		fileDebugging := zapcore.AddSync(&lumberjack.Logger{
			Filename:   "logs/application.log",
			MaxSize:    100, // MB
			MaxBackups: 3,
			MaxAge:     28, // Days
			LocalTime:  true,
		})

		fileCore := zapcore.NewCore(encoder, fileDebugging, zapcore.InfoLevel)

		// Use an AsyncWriter to write to the file asynchronously
		logger = zap.New(fileCore, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	}
	return &ZapLogger{logger: logger}
}

func (l *ZapLogger) Instance() any {
	return l.logger
}

func (l *ZapLogger) Debug(ctx context.Context, msg string, fields ...any) {
	l.logger.Debug(msg, l.convertFields(fields, ctx)...)
}

func (l *ZapLogger) Info(ctx context.Context, msg string, fields ...any) {
	l.logger.Info(msg, l.convertFields(fields, ctx)...)
}

func (l *ZapLogger) Warn(ctx context.Context, msg string, fields ...any) {
	l.logger.Warn(msg, l.convertFields(fields, ctx)...)
}

func (l *ZapLogger) Error(ctx context.Context, msg string, fields ...any) {
	l.logger.Error(msg, l.convertFields(fields, ctx)...)
}

func (l *ZapLogger) Fatal(ctx context.Context, msg string, fields ...any) {
	l.logger.Fatal(msg, l.convertFields(fields, ctx)...)
}

func (l *ZapLogger) convertFields(fields []any, ctx context.Context) []zap.Field {
	fieldsLength := len(fields)
	fieldsCount := fieldsLength / 2
	oddFields := fieldsLength%2 != 0
	if oddFields {
		fieldsCount += 1
	}
	zapFields := make([]zap.Field, 0, fieldsCount)
	values := ctx.Value(ContextKeyValues)
	if values != nil {
		for k, v := range values.(ContextValue) {
			zapFields = append(zapFields, zap.Any(string(k), v))
		}
	}
	fieldStart := 0
	if oddFields {
		fieldStart = 1
		zapFields = append(zapFields, zap.Any("message", fields[0]))
	}
	for i := fieldStart; i < len(fields); i += 2 {
		key := fields[i].(string)
		value := fields[i+1]
		zapFields = append(zapFields, zap.Any(key, value))
	}
	return zapFields
}
