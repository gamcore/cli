package logger

import "strings"

type Level int

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelPanic
)

func OfLevel(level string) Level {
	lvl := LevelInfo
	switch strings.ToLower(level) {
	case "trace":
		lvl = LevelTrace
		break
	case "debug":
		lvl = LevelDebug
		break
	case "info":
		lvl = LevelInfo
		break
	case "warn":
		lvl = LevelWarn
		break
	case "error":
		lvl = LevelError
		break
	}

	return lvl
}

func (l Level) String() string {
	lvl := "<none>"
	switch l {
	case LevelTrace:
		lvl = "trace"
		break
	case LevelDebug:
		lvl = "debug"
		break
	case LevelInfo:
		lvl = "info"
		break
	case LevelWarn:
		lvl = "warn"
		break
	case LevelError:
		lvl = "error"
		break
	case LevelPanic:
		lvl = "panic"
		break
	}

	return lvl
}
