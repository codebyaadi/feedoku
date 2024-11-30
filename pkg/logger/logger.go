package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func New() error {
	// Console encoder with color for stdout
	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "ts",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder, // Colorize log levels for the console
		EncodeCaller:  zapcore.ShortCallerEncoder,
	})

	// File encoder for plain text (non-colored)
	fileEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "ts",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeLevel:   zapcore.CapitalLevelEncoder, // No color in the file, but still show log level
		EncodeCaller:  zapcore.ShortCallerEncoder,
	})

	logFile, err := os.OpenFile("feedoku-logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	fileWriter := zapcore.AddSync(logFile)

	logLevel := zap.NewAtomicLevelAt(zap.DebugLevel)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel),
		zapcore.NewCore(fileEncoder, fileWriter, logLevel),
	)

	Logger = zap.New(core)

	return nil
}
