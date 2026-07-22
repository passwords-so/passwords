package vault

import (
	"context"
	"time"

	"github.com/novmbrs/passwords/internal/storage"
	"github.com/novmbrs/passwords/internal/utils"
	"github.com/novmbrs/passwords/internal/vaultcrypto"
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
	// Generate the vault ID before creating its password slot.
	// Ask vaultcrypto to create a random vault key wrapped by the master password.
	// Save the password slot in the new vault header.
	// Start an empty unlocked session containing the plaintext vault key.

	id := utils.GenerateID("vault")
	vaultKey, slot, err := vaultcrypto.CreatePasswordSlot(password, id)
	if err != nil {
		clear(vaultKey)
		return err
	}

	now := time.Now().UTC()
	header := storage.VaultHeader{
		Version:      slot.Version,
		ID:           id,
		Name:         name,
		PasswordSlot: slot,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := s.store.CreateHeader(ctx, header); err != nil {
		return err
	}

	s.session = &Session{
		vault: &Vault{
			ID:        id,
			Name:      name,
			Version:   1,
			CreatedAt: now,
			UpdatedAt: now,
			Items:     make(map[string]VaultItem),
		},
		vaultKey: vaultKey,
	}

	return nil
}

// Unlock loads the vault, verifies the password, and opens an in-memory session.
func (s *Service) Unlock(ctx context.Context, password []byte) error {
	// Load the vault header and ask vaultcrypto to open its password slot.
	// Treat a failed open as a generic unlock failure.
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
