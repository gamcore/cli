package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

var logger Logger

type Logger struct {
	stdout  io.Writer
	stderr  io.Writer
	logFile os.File
	level   Level
}

func New(level Level, logfile string) {
	var stdout io.Writer = os.Stdout
	var stderr io.Writer = os.Stderr
	file, _ := os.OpenFile(path.Clean(logfile), os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModeType)
	file.WriteString(fmt.Sprintf(" --- Logging Started %s --- \n", time.Now().UTC().Format(time.RFC3339)))

	logger = Logger{stdout, stderr, *file, level}
}

func (l Logger) Out() io.Writer {
	return l.stdout
}

func (l Logger) Err() io.Writer {
	return l.stderr
}

func (l Logger) LogFile() io.Writer {
	return fWriter{file: l.logFile}
}

func (l Logger) Debugf(message string, f ...any) {
	l.logf(Debug, message, f...)
}

func (l Logger) Debug(message string) {
	l.log(Debug, message)
}

func (l Logger) Infof(message string, f ...any) {
	l.logf(Info, message, f...)
}

func (l Logger) Info(message string) {
	l.log(Info, message)
}

func (l Logger) Warnf(message string, f ...any) {
	l.logf(Warn, message, f...)
}

func (l Logger) Warn(message string) {
	l.log(Warn, message)
}

func (l Logger) Errorf(message string, f ...any) {
	l.logf(Error, message, f...)
}

func (l Logger) Error(message string) {
	l.log(Error, message)
}

func (l Logger) Panicf(message string, f ...any) {
	l.logf(Panic, message, f...)
}

func (l Logger) Panic(message string) {
	l.log(Panic, message)
}

func (l Logger) Fatalf(message string, f ...any) {
	l.logf(Fatal, message, f...)
}

func (l Logger) Fatal(message string) {
	l.log(Fatal, message)
}

func (l Logger) log(level Level, message string) {
	var timestamp = time.Now().UTC().Format(time.RFC3339)
	out := fmt.Sprintf("%s [ %s ] | %s\n", timestamp, level.String(), message)
	if level >= l.level {
		wout := l.stdout
		if level >= Error {
			wout = l.stderr
		}
		if level > Error {
		}
		wout.Write([]byte(out))
	}
	l.logFile.WriteString(out)
	if level >= Panic {
		panic(out)
	}
	if level == Fatal {
		os.Exit(1)
	}
}

func (l Logger) logf(level Level, message string, f ...any) {
	l.log(level, fmt.Sprintf(message, f...))
}

type fWriter struct {
	io.Writer
	file os.File
}

func (w fWriter) Write(data []byte) (int, error) {
	return w.file.Write(data)
}
