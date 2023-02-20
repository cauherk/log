package log

import (
	"github.com/cauherk/log/logger"
	"testing"

	"go.uber.org/zap"
)

func TestLog(t *testing.T) {
	Config.SetLevel("debug")
	Debug("Debug")
	Debugf("%s", "Debuf")
	Debugw("Debugw", zap.String("Debugw", "Debugw"))
	Info("Info")
	Infof("%s", "Infof")
	Infow("Infow", zap.String("Infow", "Infow"))
	Warn("Warn")
	Warnf("%s", "Warnf")
	Warnw("Warnw", zap.String("Warnw", "Warnw"))
	Error("Error")
	Errorf("%s", "Errorf")
	Errorw("Errorw", zap.String("Errorw", "Errorw"))
	Config.DisableJSONFormat()
	Config.EnableConsoleOut()
	Config.SetProjectName("SetProjectName")
	ApplyConfig()
	Info("SetProjectName")
}

func TestLog2(t *testing.T) {
	log1 := logger.New()
	log1.Config.SetProjectName("log1")
	log1.Config.DisableJSONFormat()
	log1.ApplyConfig()
	log1.Debug("debug msg log1")
	log1.Info("info msg log1")
	log1.Warn("warn msg log1")
	log1.Error("error msg log1")
}
