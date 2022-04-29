package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stachu540/goo/internal"
	"github.com/stachu540/goo/internal/logger"
	"github.com/stachu540/goo/internal/utils"
	"os"
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
)

func init() {
	root.InitDefaultHelpCmd()
	root.InitDefaultVersionFlag()
	root.SetVersionTemplate(internal.Version)
}

func Run() {
	if err := root.Execute(); err != nil {
		logger.Errorf("%s", err)
	}
}

func doCheck(cmd *cobra.Command, argv []string) error {
	rawLevel := cmd.PersistentFlags().StringP("log", "l", "info", "Logging level [debug, info, warn, error] default: info")
	level := logger.LevelInfo
	if utils.AnySlice([]string{"debug", "info", "warn", "error"}, func(l string) bool { return *rawLevel == l }) {
		level = logger.OfLevel(*rawLevel)
	}
	debug := cmd.PersistentFlags().BoolP("debug", "d", false, "Debug mode")
	if *debug {
		level = logger.LevelDebug
	}

	err := internal.Setup(level)

	if !utils.AnySlice(argv, func(a string) bool { return a == "init" }) {
		return err
	}

	return nil
}
