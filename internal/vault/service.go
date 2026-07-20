package vault

import (
	"context"

	"github.com/novmbrs/passwords/internal/storage"
)

// Service is the core password-manager backend.
// It owns vault lifecycle operations and the unlocked session state.
type Service struct {
	store   storage.Store
	session *Session
}

// NewService wires the vault backend to a durable storage implementation.
func NewService(store storage.Store) *Service {
	return &Service{store: store}
}

// Create initializes a new encrypted vault.
func (s *Service) Create(ctx context.Context, name string, password []byte) error {
	// Create password-derivation settings and derive a key from the password.
	// Generate a random vault key and wrap it with the password-derived key.
	// Save the new vault header and start an empty unlocked session.
	return nil
}

// Unlock loads the vault, verifies the password, and opens an in-memory session.
func (s *Service) Unlock(ctx context.Context, password []byte) error {
	// Load the vault header and derive a key from the supplied password.
	// Unwrap the vault key; failure means the password is incorrect.
	// Load and decrypt every stored item into a new in-memory session.
	return nil
}

// Lock clears the unlocked in-memory session.
func (s *Service) Lock() {
	// Clear the vault key from memory.
	// Remove the unlocked session.
}

// IsUnlocked reports whether the service currently has an open session.
func (s *Service) IsUnlocked() bool {
	// Return whether an unlocked session exists.
	return false
}

// AddLogin adds a login item to the unlocked vault.
func (s *Service) AddLogin(ctx context.Context, input AddLoginInput) (VaultItem, error) {
	// Require an unlocked session and validate the input.
	// Create the login item with an ID and timestamps.
	// Encrypt and save the item, then add it to the in-memory vault.
	return VaultItem{}, nil
}

// ListItems returns safe item data from the unlocked vault.
func (s *Service) ListItems(ctx context.Context) ([]VaultItem, error) {
	// Require an unlocked session.
	// Copy the in-memory items without their password bytes.
	return nil, nil
}

// RevealPassword returns the password bytes for one unlocked vault item.
func (s *Service) RevealPassword(ctx context.Context, id string) ([]byte, error) {
	// Require an unlocked session and find the requested item.
	// Return a copy of its password bytes.
	return nil, nil
}
