package vaultcrypto

// KDFParams records how a master password is turned into a key-encryption key.
// The salt and cost settings are public and must be stored with the vault header.
type KDFParams struct {
	Algorithm string
	Salt      []byte
	MemoryKiB uint32
	Time      uint32
	Threads   uint8
	KeyLen    uint32
}

// Ciphertext is an authenticated encrypted payload and the metadata needed to open it.
type Ciphertext struct {
	Algorithm string
	Nonce     []byte
	Data      []byte
}

// PasswordSlot stores everything needed to unlock a vault key with a master password.
type PasswordSlot struct {
	Version         int
	KDF             KDFParams
	WrappedVaultKey Ciphertext
}
