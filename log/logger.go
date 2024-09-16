package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	g_logger *zap.Logger = nil

	g_mapLevel = map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}
)

// maxSize: M
// maxAge: day
func newLogger(path, filename, level string, maxSize, maxAge, maxBackups int, localtime, compress bool) *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", path, filename),
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		LocalTime:  localtime,
		Compress:   compress,
	})

	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeCaller = zapcore.FullCallerEncoder
	cfg.CallerKey = "caller"

	lvl, ok := g_mapLevel[level]
	if !ok {
		lvl = zapcore.WarnLevel
	}
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg),
			w,
			lvl,
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(cfg),
			zapcore.AddSync(os.Stdout),
			lvl,
		),
	)

	g_logger = zap.New(core)

	g_logger.Info("logger initialized.")

	return g_logger
}

// mod only interface.
func GetLogger(path, filename, level string, maxSize, maxAge, maxBackups int, localtime, compress bool) *zap.Logger {
	return newLogger(path, filename, level, maxSize, maxAge, maxBackups, localtime, compress)
}

func Close() error {
	if g_logger != nil {
		return g_logger.Sync()
	}
	return nil
}

func Debug(msg string, fields ...zapcore.Field) {
	g_logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	g_logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	g_logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	g_logger.Error(msg, fields...)
}
