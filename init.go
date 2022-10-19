// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

// Package goabu implements a distributed Event-Condition-Action engine with attribute-based interaction.
package goabu

import (
	"io"

	"github.com/abu-lang/goabu/config"

	"github.com/hyperjumptech/grule-rule-engine/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	logger.Log.Logger.Out = io.Discard
	if config.Production {
		return
	}
	zapCfg, ok := config.LogPreset("console").(zap.Config)
	if !ok {
		return
	}
	zapCfg.Level.SetLevel(zapcore.DebugLevel)
	zapLogger, err := zapCfg.Build()
	if err != nil {
		return
	}
	stdLogger, err := zap.NewStdLogAt(zapLogger, zapcore.DebugLevel)
	if err != nil {
		return
	}
	logger.Log.Logger.Out = stdLogger.Writer()
}
