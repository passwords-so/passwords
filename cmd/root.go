package cmd

import (
	"github.com/spf13/cobra"
)

/*
* $ passwords // runs the tui
* $ passwords init // create a new vault
* $ passwords unlock // unlocks the vault for some time
* $ passwords lock // locks the vault
 */

var (
	vaultPath string

	rootCmd = &cobra.Command{
		Use:          "passwords",
		Short:        "A local password manager",
		Long:         "A local, offline, secure, and open-source password manager.",
		SilenceUsage: true,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&vaultPath, "vault", "passwords.db", "path to the local vault database")
	rootCmd.MarkFlagRequired("vault")
}

func Execute() int {
	if err := rootCmd.Execute(); err != nil {
		return 1
	}
	return 0
}
