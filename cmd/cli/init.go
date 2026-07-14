package cli

import (
	"github.com/spf13/cobra"
)

func newInitCmd(newBackend func() (Backend, error)) *cobra.Command {
	return &cobra.Command{
		Use:   "init <name>",
		Short: "Create a new vault",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			backend, err := newBackend()
			if err != nil {
				return err
			}

			return runInit(cmd.Context(), backend, args)
		},
	}
}
