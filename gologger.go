package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

// calldepth is defined in log.go and uses 2 hence we use 3
const (
	calldepth = 3

	LstdFlags   = log.Ldate | log.Ltime
	LtraceFlags = log.Ldate | log.Ltime | log.Lshortfile
	LdebugFlags = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile
)

// LogLevel type
type LogLevel int

const (
	// Debug level
	Debug LogLevel = iota
	// Info level
	Info
	// Warn level
	Warn
	// Error level
	Error
	// Fatal level
	Fatal
)

var llToStr = map[LogLevel]string{
	Fatal: "[FATAL]",
	Error: "[ERROR]",
	Warn:  "[WARN]",
	Info:  "[INFO]",
	Debug: "[DEBUG]",
}

// A Logger represents an active logging object that generates lines of output to an io.Writer.
type Logger struct {
	mu     sync.Mutex // protects the following fields
	logger *log.Logger
	level  LogLevel
}

// New creates a new Logger.
func New(out io.Writer) *Logger {
	var logger *log.Logger

	if out == nil {
		logger = log.New(os.Stderr, "", LstdFlags)
	} else {
		logger = log.New(out, "", LstdFlags)
	}
	return &Logger{level: Info, logger: logger}
}

// LogLevel returns the output log level for the standard logger.
func (l *Logger) LogLevel() LogLevel {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level
}

// SetLogLevel sets the output log level for the standard logger.
func (l *Logger) SetLogLevel(ll LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if ll >= Debug && ll <= Fatal {
		l.level = ll
	}
}

// Flags returns the output flags for the standard logger.
func (l *Logger) Flags() int {
	return l.logger.Flags()
}

// SetFlags sets the output flags for the standard logger.
func (l *Logger) SetFlags(flag int) {
	l.logger.SetFlags(flag)
}

// Prefix returns the output prefix for the logger.
func (l *Logger) Prefix() string {
	return l.logger.Prefix()
}

// SetPrefix sets the output prefix for the logger.
func (l *Logger) SetPrefix(prefix string) {
	l.logger.SetPrefix(prefix)
}

// EnableStdOutput provides a shortcut for setting std flags/prefix
func (l *Logger) EnableStdOutput() {
	l.logger.SetFlags(LstdFlags)
	l.logger.SetPrefix("")
}

// EnableDebugOutput provides a shortcut for setting debug flags/prefix
func (l *Logger) EnableDebugOutput() {
	l.logger.SetFlags(LdebugFlags)
	l.logger.SetPrefix(fmt.Sprintf("[%d %s] ", os.Getpid(), os.Args[0]))
}

// EnableTraceOutput provides a shortcut for setting trace flags/prefix
func (l *Logger) EnableTraceOutput() {
	l.logger.SetFlags(LtraceFlags)
	l.logger.SetPrefix("")
}

func (l *Logger) outputln(ll LogLevel, v ...interface{}) {
	if l.level <= ll {
		v = append([]interface{}{llToStr[ll]}, v...)
		l.logger.Output(calldepth, fmt.Sprintln(v...))
	}
}

func (l *Logger) outputf(ll LogLevel, format string, v ...interface{}) {
	if l.level <= ll {
		l.logger.Output(calldepth, fmt.Sprintf(llToStr[ll]+" "+format, v...))
	}
}

func (l *Logger) Debugln(v ...interface{}) { l.outputln(Debug, v...) }

func (l *Logger) Debugf(format string, v ...interface{}) { l.outputf(Debug, format, v...) }

func (l *Logger) Infoln(v ...interface{}) { l.outputln(Info, v...) }

func (l *Logger) Infof(format string, v ...interface{}) { l.outputf(Info, format, v...) }

func (l *Logger) Warnln(v ...interface{}) { l.outputln(Warn, v...) }

func (l *Logger) Warnf(format string, v ...interface{}) { l.outputf(Warn, format, v...) }

func (l *Logger) Errorln(v ...interface{}) { l.outputln(Error, v...) }

func (l *Logger) Errorf(format string, v ...interface{}) { l.outputf(Error, format, v...) }

func (l *Logger) Fatalln(v ...interface{}) { l.outputln(Fatal, v...) }

func (l *Logger) Fatalf(format string, v ...interface{}) { l.outputf(Fatal, format, v...) }
