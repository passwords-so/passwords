package cli

import "context"

// runInit creates a new vault file.
func runInit(ctx context.Context, backend Backend, args []string) error

// runUnlock verifies the master password and opens an in-memory session.
func runUnlock(ctx context.Context, backend Backend, args []string) error

// runAddLogin adds a login item to the unlocked vault.
func runAddLogin(ctx context.Context, backend Backend, args []string) error

// runList shows item summaries from the unlocked vault.
func runList(ctx context.Context, backend Backend, args []string) error

// runReveal prints or returns the password for one item.
func runReveal(ctx context.Context, backend Backend, args []string) error

// runLock clears the unlocked session.
func runLock(ctx context.Context, backend Backend, args []string) error
