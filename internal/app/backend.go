package app

import (
	"context"

	"github.com/novmbrs/passwords/internal/vault"
)

// Backend is the CLI's narrow view of the password-manager backend.
type Backend interface {
	Create(ctx context.Context, name, pasword string) error
	Unlock(ctx context.Context, password string) error
	Lock()
	IsUnlocked() bool
	AddLogin(ctx context.Context, input vault.AddLoginInput) (vault.VaultItem, error)
	ListItems(ctx context.Context) ([]vault.VaultItem, error)
	RevealPassword(ctx context.Context, id string) ([]byte, error)
}

// NewBackend wires the CLI to SQLite storage and the vault service.
func NewBackend(vaultPath string) (Backend, error) {
	return nil, nil
}
