package log

type Logger interface {
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Warn(format string, args ...any)
	Error(format string, args ...any)
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

func Debug(format string, args ...any) {
	logger.Debug(format, args...)
}

func Info(format string, args ...any) {
	logger.Info(format, args...)
}

func Warn(format string, args ...any) {
	logger.Warn(format, args...)
}

func Error(format string, args ...any) {
	logger.Error(format, args...)
}

func Flush() {
	logger.Flush()
}
