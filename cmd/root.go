package cmd

import (
	"github.com/novmbrs/passwords/internal/config"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			// start the tui
			return nil
		},
	}
)

func init() {
	rootCmd.Flags().StringVar(&vaultPath, "vault", "", "path to the vault you want to use")
}

func Execute() int {
	if err := config.Load(); err != nil {
		return 1
	}

	if err := rootCmd.Execute(); err != nil {
		return 1
	}
	return 0
}
