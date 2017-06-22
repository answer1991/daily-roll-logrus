package drl

import (
	"github.com/Sirupsen/logrus"
	"sync"
)

var (
	logMap = make(map[string]*logrus.Logger)
	locker = &sync.Mutex{}

	enableStdout = true
	level        = logrus.DebugLevel
)

func SetEnableStdout(enable bool) {
	enableStdout = enable
}

func SetLevel(l logrus.Level) {
	locker.Lock()
	defer locker.Unlock()

	for _, theLogger := range logMap {
		theLogger.Level = l
	}

	level = l
}

func initLogger(moduleName string) {
	locker.Lock()
	defer locker.Unlock()

	_, exist := logMap[moduleName]

	if exist {
		return
	}

	logMap[moduleName] = newLogrusLogger(moduleName)
}

func GetLogger(moduleName string) *logrus.Logger {
	l, exist := logMap[moduleName]

	if exist {
		return l
	}

	initLogger(moduleName)

	l, _ = logMap[moduleName]

	return l
}
