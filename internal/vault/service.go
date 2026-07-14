package vault

import (
	"context"

	"github.com/novembersoftware/passwords/internal/storage"
)

// Service is the core password-manager backend.
// It owns vault lifecycle operations and the unlocked session state.
type Service struct {
	store   storage.Store
	session *Session
}

// NewService wires the vault backend to a durable storage implementation.
func NewService(store storage.Store) *Service {
	return nil
}

// Create initializes a new encrypted vault.
func (s *Service) Create(ctx context.Context, name string, password []byte) error {
	return nil
}

// Unlock loads the vault, verifies the password, and opens an in-memory session.
func (s *Service) Unlock(ctx context.Context, password []byte) error {
	return nil
}

// Lock clears the unlocked in-memory session.
func (s *Service) Lock() {}

// IsUnlocked reports whether the service currently has an open session.
func (s *Service) IsUnlocked() bool {
	return false
}

// AddLogin adds a login item to the unlocked vault.
func (s *Service) AddLogin(ctx context.Context, input AddLoginInput) (VaultItem, error) {
	return VaultItem{}, nil
}

// ListItems returns safe item data from the unlocked vault.
func (s *Service) ListItems(ctx context.Context) ([]VaultItem, error) {
	return nil, nil
}

// RevealPassword returns the password bytes for one unlocked vault item.
func (s *Service) RevealPassword(ctx context.Context, id string) ([]byte, error) {
	return nil, nil
}
