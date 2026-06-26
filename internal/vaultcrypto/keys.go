package vaultcrypto

// GenerateVaultKey creates the random data-encryption key for a vault.
func GenerateVaultKey() ([]byte, error)

// WrapKey encrypts the vault key with a key-encryption key.
func WrapKey(kek []byte, vaultKey []byte) (Envelope, error)

// UnwrapKey decrypts the vault key with a key-encryption key.
func UnwrapKey(kek []byte, wrapped Envelope) ([]byte, error)
