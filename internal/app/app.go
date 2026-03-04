package app

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/universtar-org/ust/internal/api"
	"github.com/universtar-org/ust/internal/log"
)

type App struct {
	Client *api.Client
	Ctx    context.Context
}

func New(token string) *App {
	return &App{
		Client: api.NewClient(token),
		Ctx:    context.Background(),
	}
}

func (a *App) RootCmd() *cobra.Command {
	var (
		token string
		debug bool
	)

	rootCmd := &cobra.Command{
		Use:   "ust [subcommand]",
		Short: "A tool used to fetch data from GitHub or update local data",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			log.InitLogger(debug)
			return nil
		},
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug log")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "GitHub token")
	rootCmd.AddCommand(a.UniqueCmd())
	rootCmd.AddCommand(a.CheckCmd())
	rootCmd.AddCommand(a.UpdateCmd())

	return rootCmd
}
