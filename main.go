package main

import (
	"os"

	"github.com/universtar-org/ust/internal/app"
)

func main() {
	app := app.New("")

	if err := app.RootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
