package main

import (
	"os"

	"github.com/novmbrs/passwords/cmd"
	"github.com/novmbrs/passwords/internal/config"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		os.Exit(1)
	}

	os.Exit(cmd.Execute())
}
