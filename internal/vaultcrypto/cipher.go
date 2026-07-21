package vaultcrypto

// seal encrypts and authenticates plaintext with AES-256-GCM.
func seal(key []byte, plaintext []byte, aad []byte) (Ciphertext, error)

// Validate the key size.
// Generate a fresh random nonce.
// Encrypt the plaintext and authenticate the associated data.
// Return the algorithm, nonce, and sealed bytes.

// open decrypts ciphertext and verifies its authentication data.
func open(key []byte, ciphertext Ciphertext, aad []byte) ([]byte, error)

// Validate the algorithm, key, nonce, and ciphertext.
// Open the AES-GCM payload using the same associated data.
// Return a generic authentication error when opening fails.
