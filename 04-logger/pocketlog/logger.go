package pocketlog

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Logger represents Logging Data structure
type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns you a logger, ready to log at the required threshold.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}
	for _, opt := range opts {
		opt(lgr)
	}
	return lgr
}

// Level represents an available logging level.
type Level byte

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}
	l.logf(LevelDebug, format, args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}
	l.logf(LevelInfo, format, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}
	l.logf(LevelError, format, args...)
}

func (l *Logger) logf(level Level, format string, args ...any) {
	var levelName string
	switch level {
	case LevelDebug:
		levelName = "DEBUG"
	case LevelInfo:
		levelName = "INFO"
	case LevelError:
		levelName = "ERROR"
	}
	fmt.Printf("%s [at %s]:  ", levelName, time.Now().Format("2006-01-02 15:04:05"))
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}
