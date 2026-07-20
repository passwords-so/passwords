package main

import (
	"os"

	"github.com/novembersoftware/passwords/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
