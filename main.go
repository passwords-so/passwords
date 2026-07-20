package main

import (
	"os"

	"github.com/novmbrs/passwords/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
