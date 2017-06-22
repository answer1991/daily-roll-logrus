package drl

import (
	"github.com/Sirupsen/logrus"
)

func newLogrusLogger(moduleName string) *logrus.Logger {
	l := logrus.New()

	l.Out = newDailyRollWriter(moduleName)
	l.Level = level

	return l
}
