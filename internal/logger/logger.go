package logger

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

type Logger struct {
	lvl      Level
	file     os.File
	format   string
	panicLvl Level
}

type Settings struct {
	Level      Level
	Format     string
	LogPath    string
	PanicLevel Level
}

func New(settings *Settings) (*Logger, error) {
	date := time.Now().UTC().Format("02-01-2006")
	file, err := os.OpenFile(path.Join(settings.LogPath, fmt.Sprintf("%s.log", date)), os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeType)
	if err != nil {
		return nil, err
	}

	return &Logger{
		lvl:      settings.Level,
		file:     *file,
		format:   settings.Format,
		panicLvl: settings.PanicLevel,
	}, nil
}

func (l *Logger) Close() error {
	return l.file.Close()
}

func (l Logger) println(level Level, message string) (int, error) {
	var writer io.Writer = &l.file
	if level >= l.lvl {
		writer = io.MultiWriter(&l.file, os.Stdout)
		if level >= LevelError {
			writer = io.MultiWriter(&l.file, os.Stderr)
		}
	}

	text, err := l.formatter(level, message)
	if err != nil {
		return 0, err
	}

	return writer.Write([]byte(*text))
}

func (l Logger) printf(level Level, message string, v ...interface{}) (int, error) {
	return l.println(level, fmt.Sprintf(message, v))
}
