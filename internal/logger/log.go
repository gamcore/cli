package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

var (
	root iLogger
)

type iLogger struct {
	defaultLevel Level
	logFile      *os.File
}

func Close() error {
	return root.logFile.Close()
}

func InitLogger(level Level, gooPath string) {
	var fout *os.File = nil
	if _, err := os.Stat(gooPath); !os.IsNotExist(err) {
		logfile := path.Join(gooPath, "goo.log")
		if _, err = os.Stat(logfile); os.IsExist(err) {
			i := 1
			for _, err = os.Stat(path.Join(gooPath, "goo.log."+string(rune(i)))); os.IsExist(err); i++ {
				_ = os.Rename(path.Join(gooPath, "goo.log."+string(rune(i))), path.Join(gooPath, "goo.log."+string(rune(i+1))))
			}
			_ = os.Rename(logfile, logfile+".1")
		}
		fout, _ = os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY, os.ModeType)
	}
	root = iLogger{level, fout}
}

func Out() io.Writer {

	if root.logFile != nil {
		f := root.logFile
		return io.MultiWriter(f, os.Stdout)
	}
	return os.Stdout
}

func Debugf(message string, f ...any) {
	logf(LevelDebug, message, f...)
}

func Debug(message string) {
	logL(LevelDebug, message)
}

func Infof(message string, f ...any) {
	logf(LevelInfo, message, f...)
}

func Info(message string) {
	logL(LevelInfo, message)
}

func Warnf(message string, f ...any) {
	logf(LevelWarn, message, f...)
}

func Warn(message string) {
	logL(LevelWarn, message)
}

func Errorf(message string, f ...any) {
	logf(LevelError, message, f...)
}

func Error(message string) {
	logL(LevelError, message)
}

func Panicf(message string, f ...any) {
	logf(LevelPanic, message, f...)
}

func Panic(message string) {
	logL(LevelPanic, message)
}

func logL(level Level, message string) {
	wout := os.Stdout
	if level >= LevelError {
		wout = os.Stderr
	}
	var logger = log.New(io.MultiWriter(root.logFile, wout), level.String(), log.LstdFlags|log.Lmsgprefix)
	logger.Println(message)
}

func logf(level Level, message string, f ...any) {
	logL(level, fmt.Sprintf(message, f...))
}
