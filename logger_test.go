package drl

import (
	"github.com/Sirupsen/logrus"
	"testing"
)

func TestGetLogger(t *testing.T) {
	logger := GetLogger("test-module")

	if nil == logger {
		t.Fail()
	}

	SetLevel(logrus.InfoLevel)
	SetEnableStdout(false)

	logger.Info("hello")
	logger.Debug("debug")
}
