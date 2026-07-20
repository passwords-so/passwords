package cmd

import (
	"github.com/novembersoftware/passwords/internal/app"
	"github.com/novembersoftware/passwords/internal/app/commands"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <name>",
	Short: "Create a new vault",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		backend, err := app.NewBackend(vaultPath)
		if err != nil {
			return err
		}

		return commands.RunInitCmd(cmd.Context(), backend, args)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
