// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LogDebug = iota - 1
	LogInfo
	LogWarning
	LogError
	LogFatal
)

type LogConfig struct {
	// "console" == "" or "json"
	Encoding string
	Level    int
}

func LogPreset(encoding string) interface{} {
	zapEnc := zapcore.EncoderConfig{
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	return zap.Config{
		Level:            zap.NewAtomicLevel(),
		Development:      !Production,
		Encoding:         encoding,
		EncoderConfig:    zapEnc,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

//----------------------------------TESTING-----------------------------------

var TestsLogConfig = LogConfig{
	Encoding: "console",
	Level:    LogDebug,
}
