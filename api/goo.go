package api

import (
	"github.com/goo-app/cli/log"
	"gopkg.in/yaml.v2"
	"os"
	"path"
	"strings"

	"github.com/hashicorp/go-version"
)

var (
	Path   string
	Config GooConfig
)

type (
	Deprecated struct {
		Reason      string  `json:"reason" yaml:"reason"`
		Since       string  `json:"since" yaml:"since"`
		Alternative *string `json:"alternative,omitempty" yaml:"alternative,omitempty"`
	}
	GooConfig struct {
		Debug           bool      `yaml:"debug"`
		LogLevel        log.Level `yaml:"log_level"`
		LogFormat       string    `yaml:"log_format"`
		LogPanic        log.Level `yaml:"log_panic"`
		FollowRedirects bool      `yaml:"follow_redirects"`
		RedirectCount   int       `yaml:"redirect_count"`
		CheckUpdates    bool      `yaml:"check_updates"`
		SkipFailure     bool      `yaml:"skip_failure"`
	}
)

func init() {
	Config = GooConfig{
		Debug:           false,
		LogLevel:        log.LevelInfo,
		FollowRedirects: true,
		RedirectCount:   10,
		CheckUpdates:    true,
		SkipFailure:     false,
		LogFormat:       "{{.Time}} | [{{.Level}}] | {{.Message}}",
		LogPanic:        log.LevelFatal,
	}
}

func Setup(debug bool, level log.Level) error {
	configPath := path.Join(Path, "config.yaml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return err
	}
	Config.Debug = debug
	Config.LogLevel = level
	logger, err := log.New(&log.Settings{
		Level:      Config.LogLevel,
		Format:     Config.LogFormat,
		LogPath:    path.Join(Path, "logs"),
		PanicLevel: Config.LogPanic,
	})
	if err != nil {
		return err
	}
	log.SetDefaultLogger(*logger)
	return nil
}

func (d *Deprecated) IsDeprecated() bool {
	v := GetVersion()
	return v.GreaterThanOrEqual(version.Must(version.NewSemver(d.Since)))
}

func (d *Deprecated) get(name string) *string {
	sb := new(strings.Builder)
	if d != nil && d.IsDeprecated() {
		sb.WriteString(name)
		sb.WriteString(" has been deprecated: \"")
		sb.WriteString(d.Reason)
		sb.WriteString("\"")
		if d.Alternative != nil {
			sb.WriteString(" Please use " + *d.Alternative)
		}
	}
	result := strings.TrimSpace(sb.String())
	if result != "" {
		return &result
	} else {
		return nil
	}
}

func tempPath() string {
	return path.Join(Path, "tmp")
}

func binaryPath() string {
	return path.Join(Path, "bin")
}
