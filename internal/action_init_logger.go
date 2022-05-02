package internal

import (
	"github.com/goo-app/cli/internal/logger"
	"github.com/spf13/viper"
	"path"
)

func initLogger() error {
	defaultRawLevel := viper.GetString("log_level")
	if Debug {
		defaultRawLevel = "debug"
	}
	defaultLevel := logger.OfLevel(defaultRawLevel)
	panicLevel := logger.OfLevel(viper.GetString("panic_on"))

	settings := logger.Settings{Level: defaultLevel, Format: viper.GetString("log_format"), LogPath: path.Join(Path, "log"), PanicLevel: panicLevel}
	l, err := logger.New(&settings)
	if err != nil {
		return err
	}
	logger.SetDefaultLogger(*l)
	return nil
}
