package storage

import "context"

// FileStore persists a vault to one local file.
type FileStore struct {
	path string
}

// NewFileStore creates a storage adapter for a local vault file.
func NewFileStore(path string) *FileStore

// LoadHeader reads vault metadata from the local file.
func (s *FileStore) LoadHeader(ctx context.Context) (VaultHeader, error)

// SaveHeader writes vault metadata to the local file.
func (s *FileStore) SaveHeader(ctx context.Context, header VaultHeader) error

// ListItems reads all encrypted item records from the local file.
func (s *FileStore) ListItems(ctx context.Context) ([]EncryptedItemRecord, error)

// GetItem reads one encrypted item record from the local file.
func (s *FileStore) GetItem(ctx context.Context, id string) (EncryptedItemRecord, error)

// PutItem writes one encrypted item record to the local file.
func (s *FileStore) PutItem(ctx context.Context, item EncryptedItemRecord) error

// DeleteItem removes one encrypted item record from the local file.
func (s *FileStore) DeleteItem(ctx context.Context, id string) error
