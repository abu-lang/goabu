package steel

import (
	"io"
	"steel/config"

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
