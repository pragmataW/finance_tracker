package pkg

import (
	"io"
	"log"
	"sync"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

var (
	once   sync.Once
	logger *Logger
)

type Logger struct {
	logLevel      LogLevel
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	debugLogger   *log.Logger
}

func NewLogger(level LogLevel, out io.Writer) *Logger {
	once.Do(func() {
		logger = &Logger{
			logLevel:      level,
			infoLogger:    log.New(out, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
			warningLogger: log.New(out, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
			errorLogger:   log.New(out, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
			debugLogger:   log.New(out, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		}
	})
	return logger
}

func (l *Logger) Info(message string) {
	if l.logLevel <= INFO {
		l.infoLogger.Println(message)
	}
}

func (l *Logger) Warning(message string) {
	if l.logLevel <= WARNING {
		l.warningLogger.Println(message)
	}
}

func (l *Logger) Error(message string) {
	if l.logLevel <= ERROR {
		l.errorLogger.Println(message)
	}
}

func (l *Logger) Debug(message string) {
	if l.logLevel <= DEBUG {
		l.debugLogger.Println(message)
	}
}
