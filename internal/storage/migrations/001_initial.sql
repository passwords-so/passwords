CREATE TABLE vault_header (
    singleton INTEGER PRIMARY KEY CHECK (singleton = 1),
    vault_id TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    version INTEGER NOT NULL,
    kdf_algorithm TEXT NOT NULL,
    kdf_salt BLOB NOT NULL,
    kdf_memory_kib INTEGER NOT NULL,
    kdf_time INTEGER NOT NULL,
    kdf_threads INTEGER NOT NULL,
    kdf_key_len INTEGER NOT NULL,
    wrapped_key_algorithm TEXT NOT NULL,
    wrapped_key_nonce BLOB NOT NULL,
    wrapped_key_data BLOB NOT NULL,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL
) STRICT;

CREATE TABLE encrypted_items (
    id TEXT PRIMARY KEY,
    kind TEXT NOT NULL,
    version INTEGER NOT NULL,
    cipher_algorithm TEXT NOT NULL,
    cipher_nonce BLOB NOT NULL,
    cipher_data BLOB NOT NULL,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL
) STRICT;

CREATE INDEX encrypted_items_created_at_idx
    ON encrypted_items (created_at, id);
