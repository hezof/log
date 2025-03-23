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
var _logger Logger = StdoutLogger

var StdoutLogger, _ = NewFileLogger(&FileConfig{
	File: STDOUT,
})

func InitLogger(lgr Logger) {
	if _logger != nil {
		_logger.Flush()
	}
	_logger = lgr
}

func Debug(format string, args ...any) {
	_logger.Debug(format, args...)
}

func Info(format string, args ...any) {
	_logger.Info(format, args...)
}

func Warn(format string, args ...any) {
	_logger.Warn(format, args...)
}

func Error(format string, args ...any) {
	_logger.Error(format, args...)
}

func Flush() {
	_logger.Flush()
}
