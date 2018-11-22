package log

import (
	"fmt"
	"io"
	stdLog "log"
	"os"
	"strings"
)

const (
	Debug = iota
	Info
	Warn
	Error
	Fatal
	// disable log
	Off
)

type Logger struct {
	Level  int
	Logger *stdLog.Logger
}

var loggerLevel = Debug

var loggers []*Logger

func NewLogger(out io.Writer) *Logger {
	newLogger := &Logger{
		Level:  loggerLevel,
		Logger: stdLog.New(out, "", stdLog.Ldate|stdLog.Ltime|stdLog.Lshortfile),
	}

	loggers = append(loggers, newLogger)
	return newLogger
}

// set logger level of all loggers
func SetLevel(level string) {
	loggerLevel = getLevel(level)

	for _, l := range loggers {
		l.SetLevel(level)
	}
}

func getLevel(level string) int {
	level = strings.ToLower(level)

	switch level {
	case "off":
		return Off
	case "debug":
		return Debug
	case "info":
		return Info
	case "warn":
		return Warn
	case "error":
		return Error
	case "fatal":
		return Fatal
	default:
		return Info
	}
}

func (l *Logger) SetLevel(level string) {
	l.Level = getLevel(level)
}

func (l *Logger) Debug(v ...interface{}) {
	if Debug < l.Level {
		return
	}

	l.Logger.SetPrefix("[Debug]")
	l.Logger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, content ...interface{}) {
	if Debug < l.Level {
		return
	}

	l.Logger.SetPrefix("[Debug]")
	l.Logger.Output(2, fmt.Sprintf(format, content...))
}

func (l *Logger) Warn(v ...interface{}) {
	if Warn < l.Level {
		return
	}

	l.Logger.SetPrefix("[Warning]")
	l.Logger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if Warn < l.Level {
		return
	}

	l.Logger.SetPrefix("[Warning]")
	l.Logger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	if Info < l.Level {
		return
	}

	l.Logger.SetPrefix("[Info]")
	l.Logger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if Info < l.Level {
		return
	}

	l.Logger.SetPrefix("[Info]")
	l.Logger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(content ...interface{}) {
	if Fatal < l.Level {
		return
	}
	l.Logger.SetPrefix("[Fatal]")
	l.Logger.Output(2, fmt.Sprint(content...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, content ...interface{}) {
	if Fatal < l.Level {
		return
	}
	l.Logger.SetPrefix("[Fatal]")
	l.Logger.Output(2, fmt.Sprintf(format, content...))
	os.Exit(1)
}
