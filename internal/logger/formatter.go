package logger

import (
	"strings"
	"text/template"
	"time"
)

type f struct {
	Message string
	Time    string
	Level   Level
}

func (l Logger) formatter(level Level, message string) (*string, error) {
	t := template.New("logger")
	t, err := t.Parse(l.format)
	if err != nil {
		return nil, err
	}
	out := strings.Builder{}
	err = t.Execute(&out, f{message, time.Now().UTC().Format(time.RFC3339), level})
	if err != nil {
		return nil, err
	}
	raw := out.String()
	if !strings.HasSuffix(raw, "\n") {
		raw = raw + "\n"
	}
	return &raw, nil
}
