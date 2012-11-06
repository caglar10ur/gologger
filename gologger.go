package logger

import "fmt"
import "io"
import "log"
import "os"
import "sync"

const calldepth     = 3

const flags         = log.Ldate | log.Ltime
const trace_flags   = log.Ldate | log.Ltime | log.Lshortfile
const debug_flags   = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile

type LogLevel int
const (
    Debug LogLevel = iota
    Info
    Warn
    Error
    Fatal
)

var llToStr map[LogLevel]string = map[LogLevel]string {
    Fatal: "[FATAL]",
    Error: "[ERROR]",
    Warn:  "[WARN]",
    Info:  "[INFO]",
    Debug: "[DEBUG]",
}

type Logger struct {
    // ensures atomic writes
    mu      sync.Mutex
    Logger  *log.Logger
    Level   LogLevel
}

func New(out io.Writer) *Logger {
    var logger *log.Logger

    if out == nil {
        logger = log.New(os.Stderr, "", flags)
    } else {
        logger = log.New(out, "", flags)
    }
    return &Logger{Level: Info, Logger: logger}
}

func (l *Logger) SetLogLevel(ll LogLevel) {
    if ll >= Debug && ll <= Fatal {
        l.mu.Lock()
        defer l.mu.Unlock()
        l.Level = ll
    }
}

func (l *Logger) SetFlags() { l.mu.Lock(); defer l.mu.Unlock(); l.Logger.SetFlags(flags); l.Logger.SetPrefix("") }

func (l *Logger) SetTraceFlags() { l.mu.Lock(); defer l.mu.Unlock(); l.Logger.SetFlags(trace_flags); l.Logger.SetPrefix("") }

func (l *Logger) SetDebugFlags() { l.mu.Lock(); defer l.mu.Unlock(); l.Logger.SetFlags(debug_flags); l.Logger.SetPrefix(fmt.Sprintf("[%d %s] ", os.Getpid(), os.Args[0])) }

func (l *Logger) outputln(ll LogLevel, v ...interface{}) {
    if l.Level <= ll {
        v = append([]interface{}{llToStr[ll]}, v...)
        l.Logger.Output(calldepth, fmt.Sprintln(v...))
    }
}

func (l *Logger) outputf(ll LogLevel, format string, v ...interface{}) {
    if l.Level <= ll {
        l.Logger.Output(calldepth, fmt.Sprintf(llToStr[ll] + " " + format, v...))
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
