package cmd

import (
	"github.com/spf13/cobra"
	myapp "github.com/universtar-org/tools/internal/app"
	"github.com/universtar-org/tools/internal/cli/unique"
	"github.com/universtar-org/tools/internal/log"
)

var (
	token string
	debug bool
)

var app *myapp.App

var rootCmd = &cobra.Command{
	Use:   "ust [subcommand]",
	Short: "A tool used to fetch data from GitHub or update local data",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		log.InitLogger(debug)
		app = myapp.New(token)
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	getApp := func() *myapp.App { return app }

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug log")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "GitHub token")
	rootCmd.AddCommand(unique.NewCommand(getApp))
}

func Exec() {
	rootCmd.Execute()
}
