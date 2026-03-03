package main

import (
	"os"

	"github.com/universtar-org/tools/internal/app"
)

func main() {
	app := app.New("")

	if err := app.RootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
