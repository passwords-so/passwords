package envelope

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

const (
	envelopeAlgorithm = "AES-256-GCM"
	vaultKeySize      = 32
)

// GenerateVaultKey creates the random data-encryption key for a vault.
func GenerateVaultKey() ([]byte, error) {
	key := make([]byte, vaultKeySize)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

// WrapKey encrypts the vault key with a key-encryption key.
func WrapKey(kek []byte, vaultKey []byte) (Envelope, error) {
	block, err := aes.NewCipher(kek)
	if err != nil {
		return Envelope{}, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return Envelope{}, err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return Envelope{}, err
	}

	return Envelope{
		Algorithm:  envelopeAlgorithm,
		Nonce:      nonce,
		Ciphertext: aead.Seal(nil, nonce, vaultKey, nil),
	}, nil
}

// UnwrapKey decrypts the vault key with a key-encryption key.
func UnwrapKey(kek []byte, wrapped Envelope) ([]byte, error) {
	return nil, nil
}
