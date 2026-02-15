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
		"debug":   zapcore.DebugLevel,
		"info":    zapcore.InfoLevel,
		"warn":    zapcore.WarnLevel,
		"warning": zapcore.WarnLevel,
		"error":   zapcore.ErrorLevel,
	}
)

/*
path: logfile path
filename: log filename in @path. should not "error.log" which is for @errorFile
level: one of debug/info/warn/warning/error
maxSize: M
maxAge: day
maxBackups: number of most backups
localtime: use local time
compress: backups if compressed
console: enable/disable console log.(DEBUG level)
errorFile: save ERROR level logs to standalone `error.log`. meanwhile save to `main` log file too.
*/
func newLogger(path, filename, level string, maxSize, maxAge, maxBackups int, localtime, compress, console, errorFile bool) *zap.Logger {
	writeSyncerAll := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", path, filename),
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		LocalTime:  localtime,
		Compress:   compress,
	})
	writeSyncerError := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/error.log", path),
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		LocalTime:  localtime,
		Compress:   compress,
	})

	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeCaller = zapcore.FullCallerEncoder
	cfg.CallerKey = "fn"

	lvl, ok := g_mapLevel[level]
	if !ok {
		lvl = zapcore.WarnLevel
	}
	//
	cores := []zapcore.Core{
		zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg),
			writeSyncerAll,
			lvl,
		),
	}
	if console {
		cores = append(cores, zapcore.NewCore(
			zapcore.NewConsoleEncoder(cfg),
			zapcore.AddSync(os.Stdout),
			zapcore.DebugLevel,
		))
	}
	if errorFile {
		cores = append(cores, zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg),
			writeSyncerError,
			zapcore.ErrorLevel,
		))
	}

	core := zapcore.NewTee(cores...)

	g_logger = zap.New(core)

	g_logger.Info("logger initialized.")

	return g_logger
}

// mod only interface.
// @maxSize: M @maxAge days @maxBackup fileNums
func GetLogger(path, filename, level string, maxSize, maxAge, maxBackups int, localtime, compress, console, errorFile bool) *zap.Logger {
	return newLogger(path, filename, level, maxSize, maxAge, maxBackups, localtime, compress, console, errorFile)
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
