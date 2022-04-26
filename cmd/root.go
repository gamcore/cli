package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stachu540/goo/internal"
	"github.com/stachu540/goo/internal/log"
	"github.com/stachu540/goo/internal/utils"
)

var (
	root = &cobra.Command{
		Use:   "goo",
		Short: "Goo App Manager",
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
		Long:    `Manage your developer applications with cross-platform application installer`,
		PreRunE: doCheck,
	}
)

func init() {
	root.InitDefaultHelpCmd()
	root.InitDefaultVersionFlag()
	root.SetVersionTemplate(internal.Version)
}

func Run() {
	if err := root.Execute(); err != nil {
		internal.Logger.Fatalf("%s", err)
	}
}

func doCheck(cmd *cobra.Command, argv []string) error {
	rawLevel := cmd.PersistentFlags().StringP("log", "l", "info", "Logging level [debug, info, warn, error] default: info")
	level := log.Info
	if utils.AnySlice([]string{"debug", "info", "warn", "error"}, func(l string) bool { return *rawLevel == l }) {
		level = log.OfLevel(*rawLevel)
	}
	debug := cmd.PersistentFlags().BoolP("debug", "d", false, "Debug mode")
	if *debug {
		level = log.Debug
	}

	err := internal.Setup(level)

	if !utils.AnySlice(argv, func(a string) bool { return a == "init" }) {
		return err
	}

	return nil
}
