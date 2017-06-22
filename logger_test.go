package drl

import "testing"

func TestGetLogger(t *testing.T) {
	logger := GetLogger("test-module")

	if nil == logger {
		t.Fail()
	}

	logger.Info("hello world")
}
