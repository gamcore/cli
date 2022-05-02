package logger

import (
	"strings"
)

type Level int

const (
	LevelTrace Level = iota - 1
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelCritical
	LevelFatal
)

var prefixLevel = map[Level]string{
	LevelTrace:    "TRACE",
	LevelDebug:    "DEBUG",
	LevelInfo:     "INFO",
	LevelWarn:     "WARN",
	LevelError:    "ERROR",
	LevelCritical: "CRITICAL",
	LevelFatal:    "FATAL",
}

var rawPrefixLevel = map[string]Level{
	"trace":    LevelTrace,
	"debug":    LevelDebug,
	"info":     LevelInfo,
	"warn":     LevelWarn,
	"error":    LevelError,
	"critical": LevelCritical,
	"fatal":    LevelFatal,
}

func (l Level) MarshalYAML() (interface{}, error) {
	return l.String(), nil
}

func (l Level) MarshalJSON() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l *Level) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var o *string
	if err := unmarshal(&o); err != nil {
		return err
	}
	if o != nil {
		*l = OfLevel(*o)
	}
	return nil
}

func (l *Level) UnmarshalJSON(data []byte) error {
	*l = OfLevel(string(data))
	return nil
}

func OfLevel(level string) Level {
	return rawPrefixLevel[strings.ToLower(level)]
}

func (l Level) String() string {
	return strings.ToLower(prefixLevel[l])
}

func (l Level) StringC() string {
	return prefixLevel[l]
}
