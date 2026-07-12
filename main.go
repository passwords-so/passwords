package main

import (
	"os"

	"github.com/novembersoftware/passwords/cmd/cli"
)

func main() {
	os.Exit(cli.Execute())
}
