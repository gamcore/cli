package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

type Logger struct {
	lvl      Level
	file     *os.File
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
		file:     file,
		format:   settings.Format,
		panicLvl: settings.PanicLevel,
	}, nil
}

func (l *Logger) Close() error {
	return l.file.Close()
}

func (l Logger) println(level Level, message string) {
	text, _ := l.formatter(level, message)
	var writer *io.Writer
	fOut := l.file
	*writer = fOut
	if level >= l.lvl {
		out := os.Stdout
		if level >= LevelError {
			out = os.Stderr
		}
		*writer = io.MultiWriter(fOut, out)
	}

	if writer != nil {
		w := *writer
		_, _ = w.Write([]byte(*text))
	}

	if level >= l.panicLvl {
		os.Exit(1)
	}
}

func (l Logger) printf(level Level, message string, v ...interface{}) {
	l.println(level, fmt.Sprintf(message, v...))
}
