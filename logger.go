package log

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	ErrorStack(format string, args ...interface{})
	Flush()
}

type FileLogger interface {
	Logger
	Close()
}

// 默认stdout
var logger Logger

func init() {
	logger, _ = NewFileLogger(&FileConfig{
		File: STDOUT,
	})
}

func InitLogger(lgr Logger) {
	if logger != nil {
		logger.Flush()
	}
	logger = lgr
}

func Debug(format string, args ...interface{}) {
	logger.Debug(format, args...)
}

func Info(format string, args ...interface{}) {
	logger.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	logger.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	logger.Error(format, args...)
}

func ErrorStack(format string, args ...interface{}) {
	logger.ErrorStack(format, args...)
}
func Flush() {
	logger.Flush()
}
