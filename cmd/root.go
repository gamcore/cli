package cmd

import (
	"os"
	"strings"

	"github.com/goo-app-manager/goo/api"
	"github.com/goo-app-manager/goo/log"
	"github.com/goo-app-manager/goo/utils"
	"github.com/spf13/cobra"
)

var (
	root = &cobra.Command{
		Use:     "goo",
		Short:   "Goo App Manager",
		Version: api.Version,
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
		Long:              `Manage your developer applications with cross-platform application installer`,
		PersistentPreRunE: doCheck,
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			err := log.Close()

			if err == os.ErrInvalid {
				return nil
			}

			return err
		},
	}
	rawLogLevel string
	debug       bool
)

func init() {
	root.SilenceErrors = false
	root.InitDefaultHelpCmd()
	root.InitDefaultVersionFlag()
	root.SetVersionTemplate(api.Version)
	root.PersistentFlags().StringVarP(&rawLogLevel, "log", "l", "info", "Logging level [debug, info, warn, error, none]")
	root.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug mode")

}

func Run() {
	if err := root.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}

func doCheck(_ *cobra.Command, argv []string) error {
	if len(argv) >= 1 {
		if !utils.AnySlice([]string{"help", "init"}, func(c string) bool { return strings.ToLower(argv[0]) == c }) {
			if !utils.AnySlice([]string{"debug", "info", "warn", "error", "critical", "none"}, func(l string) bool { return strings.ToLower(rawLogLevel) == l }) {
				rawLogLevel = "info"
			}

			return api.Setup(debug, log.OfLevel(rawLogLevel))
		}
	}

	return nil
}
