package logger

import "log"
import "os"
import "sync"

type LogLevel int

const (
    Debug LogLevel = iota
    Info
    Warn
    Error
    Fatal
)

var logLevelToString map[LogLevel]string = map[LogLevel]string {
    Fatal: "[FATAL]",
    Error: "[ERROR]",
    Warn:  "[WARN]",
    Info:  "[INFO]",
    Debug: "[DEBUG]",
}

type Logger struct {
    // ensures atomic writes
    mu     sync.Mutex
    Level LogLevel
}

func New() *Logger {
    return &Logger{Level: Info}
}

func (l *Logger) SetLogLevel(ll LogLevel) {
    if ll >= Debug && ll <= Fatal {
        l.mu.Lock()
        defer l.mu.Unlock()
        l.Level = ll
    }
}

func (l *Logger) SetOutput(filename string) {
    logfile, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0644)
    if err != nil {
        log.Fatal("Unable to open log file: ", err)
    }
    log.SetOutput(logfile)
}

func (l *Logger) SetFlags() {
    log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}

func (l *Logger) UnsetFlags() {
    log.SetFlags(log.Ldate | log.Ltime)
}

func (l *Logger) outputln(ll LogLevel, v ...interface{}) {
    if l.Level <= ll {
        v = append([]interface{}{logLevelToString[ll]}, v...)
        log.Println(v...)
    }
}

func (l *Logger) outputf(ll LogLevel, format string, v ...interface{}) {
    if l.Level <= ll {
        log.Printf(logLevelToString[ll] + " " + format, v...)
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
