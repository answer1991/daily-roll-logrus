package drl

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

var (
	logDir = os.Getenv("DAILY_ROLL_LOGRUS_LOG_PATH")

	defaultLogPath = filepath.Join(os.Getenv("HOME"), "logs")
)

func init() {
	if "" == logDir {
		logDir = defaultLogPath
	}

	_, err := os.Stat(logDir)

	if os.IsNotExist(err) {
		err = os.MkdirAll(logDir, os.ModePerm)

		if nil != err {
			panic(err)
		}
	}
}

func newDailyRollWriter(prefixFileName string) *dailyRollWriter {
	ret := &dailyRollWriter{
		prefixFileName: prefixFileName,
		locker:         &sync.Mutex{},
	}

	runtime.SetFinalizer(ret, writerFinalizer)

	return ret
}

type dailyRollWriter struct {
	prefixFileName string
	current        string
	stdout         bool
	writer         *os.File
	locker         sync.Locker
}

func (w *dailyRollWriter) initWriter() {
	w.locker.Lock()
	defer w.locker.Unlock()

	writerFinalizer(w)

	logFile := filepath.Join(logDir, fmt.Sprintf("%s-%s.log", w.prefixFileName, w.current))

	log, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)

	if nil != err {
		if os.IsNotExist(err) {
			log, err = os.Create(logFile)
		}

		if nil != err {
			panic(err)
		}
	}

	w.writer = log
}

func (w *dailyRollWriter) Write(p []byte) (n int, err error) {
	now := time.Now().Format("2006-01-02")

	if now != w.current {
		w.current = now

		w.initWriter()
	}

	if enableStdout {
		os.Stdout.Write(p)
	}

	return w.writer.Write(p)
}

func writerFinalizer(w *dailyRollWriter) {
	if nil != w.writer {
		w.writer.Close()
	}
}
