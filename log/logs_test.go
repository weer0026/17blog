package log

import (
	"os"
	"testing"
)

var logger = NewLogger(os.Stdout)

func TestSetLevel(t *testing.T) {
	SetLevel("trace")
}

func TestDebug(t *testing.T) {
	logger.SetLevel("debug")
	logger.Debug("test Debug log")
	logger.SetLevel("off")
	logger.Debug("off log")
}

func TestInfo(t *testing.T) {
	logger.SetLevel("info")
	logger.Info("test Info log")
	logger.SetLevel("off")
	logger.Info("off log")
}