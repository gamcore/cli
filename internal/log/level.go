package log

import "strings"

type Level int

const (
	Trace Level = iota
	Debug
	Info
	Warn
	Error
	Panic
	Fatal
)

func OfLevel(level string) Level {
	lvl := Info
	switch strings.ToLower(level) {
	case "trace":
		lvl = Trace
		break
	case "debug":
		lvl = Debug
		break
	case "info":
		lvl = Info
		break
	case "warn":
		lvl = Warn
		break
	case "error":
		lvl = Error
		break
	}

	return lvl
}

func (l Level) String() string {
	lvl := "<none>"
	switch l {
	case Trace:
		lvl = "trace"
		break
	case Debug:
		lvl = "debug"
		break
	case Info:
		lvl = "info"
		break
	case Warn:
		lvl = "warn"
		break
	case Error:
		lvl = "error"
		break
	case Panic:
		lvl = "panic"
		break
	case Fatal:
		lvl = "fatal"
	}

	return lvl
}
