package cmd

import (
	"github.com/goo-app/cli/internal"
	"github.com/goo-app/cli/internal/logger"
	"github.com/goo-app/cli/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	root = &cobra.Command{
		Use:   "goo",
		Short: "Goo App Manager",
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
		Long:              `Manage your developer applications with cross-platform application installer`,
		PersistentPreRunE: doCheck,
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			err := logger.Close()

			if err == os.ErrInvalid {
				return nil
			}

			return err
		},
	}
	rawLogLevel *string
	isDebug     = false
)

func init() {
	root.SilenceErrors = false
	root.InitDefaultHelpCmd()
	root.InitDefaultVersionFlag()
	root.SetVersionTemplate(internal.Version)
	root.PersistentFlags().StringVarP(rawLogLevel, "log", "l", "info", "Logging level [debug, info, warn, error, none] default: info")
	root.PersistentFlags().BoolVarP(&isDebug, "debug", "d", false, "Debug mode")

	viper.BindPFlag("debug", root.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("log_level", root.PersistentFlags().Lookup("log"))
	viper.SetDefault("log_level", "info")
	viper.SetDefault("log_format", "{{ .Time }} [{{ .Level }}] {{ .Message }}")
	viper.SetDefault("panic_on", "fatal")
}

func Run() {
	if err := root.Execute(); err != nil {
		logger.FatalF("%s", err)
	}
}

func doCheck(_ *cobra.Command, argv []string) error {
	if !utils.AnySlice([]string{"debug", "info", "warn", "error", "critical", "none"}, func(l string) bool { return strings.ToLower(*rawLogLevel) == l }) {
		rawLogLevel = nil
	}
	viper.Set("log_level", rawLogLevel)

	internal.Debug = isDebug

	err := internal.Setup()

	if len(argv) >= 1 {
		if argv[0] != "init" {
			return err
		}
	}

	return nil
}
