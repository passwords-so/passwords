package vaultcrypto

// EncryptItem encrypts a serialized vault item with the unlocked vault key.
func EncryptItem(vaultKey []byte, vaultID string, itemID string, plaintext []byte) (Ciphertext, error)

// Build stable associated data from the vault ID and item ID.
// Seal the serialized item with the vault key.

// DecryptItem decrypts and authenticates a serialized vault item.
func DecryptItem(vaultKey []byte, vaultID string, itemID string, ciphertext Ciphertext) ([]byte, error)

// Rebuild the same associated data from the vault ID and item ID.
// Open the item with the vault key.
