package cli

import "github.com/spf13/cobra"

type BackendFactory func(vaultPath string) (Backend, error)

func NewRootCmd(factory BackendFactory) *cobra.Command {
	var vaultPath string

	root := &cobra.Command{
		Use:          "passwords",
		Short:        "A local password manager",
		Long:         "A local, offline, secure, and open-source password manager.",
		SilenceUsage: true,
	}

	root.PersistentFlags().StringVar(
		&vaultPath,
		"vault",
		"passwords.db",
		"path to the local vault database",
	)

	root.AddCommand(newInitCmd(func() (Backend, error) {
		return NewBackend(vaultPath)
	}))

	return root
}

func Execute() int {
	if err := NewRootCmd(NewBackend).Execute(); err != nil {
		return 1
	}
	return 0
}
