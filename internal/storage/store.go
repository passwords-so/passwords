package storage

import (
	"context"
	"errors"
	"time"

	"github.com/novmbrs/passwords/internal/vaultcrypto"
)

var (
	ErrVaultNotInitialized     = errors.New("vault not initialized")
	ErrVaultAlreadyInitialized = errors.New("vault already initialized")
	ErrItemNotFound            = errors.New("item not found")
)

// Store is the persistence boundary used by the vault backend.
type Store interface {
	LoadHeader(ctx context.Context) (VaultHeader, error)
	CreateHeader(ctx context.Context, header VaultHeader) error
	ListItems(ctx context.Context) ([]EncryptedItemRecord, error)
	GetItem(ctx context.Context, id string) (EncryptedItemRecord, error)
	PutItem(ctx context.Context, item EncryptedItemRecord) error
	DeleteItem(ctx context.Context, id string) error
}

// VaultHeader is safe metadata plus the encrypted vault key.
type VaultHeader struct {
	ID           string
	Name         string
	Version      int
	PasswordSlot vaultcrypto.PasswordSlot
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// EncryptedItemRecord is one encrypted vault item as stored on disk.
type EncryptedItemRecord struct {
	ID         string
	Kind       string
	Version    int
	Ciphertext vaultcrypto.Ciphertext
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
