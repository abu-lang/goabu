// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

// Package goabu implements a distributed Event-Condition-Action engine with attribute-based interaction.
package goabu

import (
	"github.com/abu-lang/goabu/config"

	"github.com/hyperjumptech/grule-rule-engine/antlr"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	var zapLogger *zap.Logger
	if config.Production {
		zapLogger = zap.NewNop()
	} else {
		zapCfg, ok := config.LogPreset("console").(zap.Config)
		if !ok {
			return
		}
		zapCfg.Level.SetLevel(zapcore.InfoLevel)
		var err error
		zapLogger, err = zapCfg.Build()
		if err != nil {
			return
		}
	}
	antlr.SetLogger(zapLogger)
	ast.SetLogger(zapLogger)
	logger.SetLogger(zapLogger)
}
