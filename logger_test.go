package log

import (
	"os"
	"testing"
	"time"
)

var stdoutLogger Logger

func init() {
	stdoutLogger, _ = NewFileLogger(&FileConfig{
		File:             STDOUT,
		DiscardThreshold: 100,
	})
}

var (
	threads = 5
	times   = 1000
)

func TestFileLogger(t *testing.T) {
	defer stdoutLogger.Flush()

	stdoutLogger.Error("这是一个错误 %v", os.ErrClosed)
	time.Sleep(time.Second)
}

func BenchmarkFileLogger(t *testing.B) {
	stdoutLogger.Error("这是一个错误 %v", os.ErrClosed)
}
