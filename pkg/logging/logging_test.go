package logging_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/lwinmgmg/user-go/pkg/logging"
)

func TestNew(t *testing.T) {
	logPath := "log.log"
	logger := logging.New(slog.LevelDebug, false, logPath)
	defer os.Remove(logPath)
	logger.Debug("Hello")
	logger = logging.New(slog.LevelInfo, true)
	logger.Info("Hello1")
	logger.Infof("Hello%v", 2)
}
