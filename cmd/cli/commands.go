package cli

import (
	"context"

	"github.com/novembersoftware/passwords/internal/cli"
)

// runInit creates a new vault file.
func runInit(ctx context.Context, backend Backend, args []string) error {
	name, password, err := cli.InitVaultForm()
	if err != nil {
		return err
	}
	if err = backend.Create(ctx, name, password); err != nil {
		return err
	}
	return nil
}

// runUnlock verifies the master password and opens an in-memory session.
func runUnlock(ctx context.Context, backend Backend, args []string) error {
	return nil
}

// runAddLogin adds a login item to the unlocked vault.
func runAddLogin(ctx context.Context, backend Backend, args []string) error {
	return nil
}

// runList shows item summaries from the unlocked vault.
func runList(ctx context.Context, backend Backend, args []string) error {
	return nil
}

// runReveal prints or returns the password for one item.
func runReveal(ctx context.Context, backend Backend, args []string) error {
	return nil
}

// runLock clears the unlocked session.
func runLock(ctx context.Context, backend Backend, args []string) error {
	return nil
}
