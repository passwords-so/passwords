package storage

import (
	"context"
	"database/sql"
)

// SQLiteStore persists encrypted vault records in a local SQLite database.
type SQLiteStore struct {
	path string
	db   *sql.DB
}

// NewSQLiteStore opens or creates a local SQLite vault database.
func NewSQLiteStore(ctx context.Context, path string) (*SQLiteStore, error) {
	return nil, nil
}

// Migrate creates or updates the SQLite schema used by the vault.
func (s *SQLiteStore) Migrate(ctx context.Context) error {
	return nil
}

// Close releases the SQLite database handle.
func (s *SQLiteStore) Close() error {
	return nil
}

// LoadHeader reads vault metadata from SQLite.
func (s *SQLiteStore) LoadHeader(ctx context.Context) (VaultHeader, error) {
	return VaultHeader{}, nil
}

// SaveHeader writes vault metadata to SQLite.
func (s *SQLiteStore) SaveHeader(ctx context.Context, header VaultHeader) error {
	return nil
}

// ListItems reads all encrypted item records from SQLite.
func (s *SQLiteStore) ListItems(ctx context.Context) ([]EncryptedItemRecord, error) {
	return nil, nil
}

// GetItem reads one encrypted item record from SQLite.
func (s *SQLiteStore) GetItem(ctx context.Context, id string) (EncryptedItemRecord, error) {
	return EncryptedItemRecord{}, nil
}

// PutItem inserts or updates one encrypted item record in SQLite.
func (s *SQLiteStore) PutItem(ctx context.Context, item EncryptedItemRecord) error {
	return nil
}

// DeleteItem removes one encrypted item record from SQLite.
func (s *SQLiteStore) DeleteItem(ctx context.Context, id string) error {
	return nil
}
