package envelope

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"testing"
)

func TestGenerateVaultKey(t *testing.T) {
	key, err := GenerateVaultKey()
	if err != nil {
		t.Fatalf("GenerateVaultKey() error = %v", err)
	}

	const wantLen = 32
	if len(key) != wantLen {
		t.Fatalf("GenerateVaultKey() length = %d, want %d", len(key), wantLen)
	}
}

func TestWrapKey(t *testing.T) {
	kek := bytes.Repeat([]byte{1}, vaultKeySize)
	vaultKey := bytes.Repeat([]byte{2}, vaultKeySize)

	wrapped, err := WrapKey(kek, vaultKey)
	if err != nil {
		t.Fatalf("WrapKey() error = %v", err)
	}

	if wrapped.Algorithm != envelopeAlgorithm {
		t.Fatalf("WrapKey() algorithm = %q, want %q", wrapped.Algorithm, envelopeAlgorithm)
	}

	block, err := aes.NewCipher(kek)
	if err != nil {
		t.Fatalf("aes.NewCipher() error = %v", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		t.Fatalf("cipher.NewGCM() error = %v", err)
	}

	if len(wrapped.Nonce) != aead.NonceSize() {
		t.Fatalf("WrapKey() nonce length = %d, want %d", len(wrapped.Nonce), aead.NonceSize())
	}
	if bytes.Equal(wrapped.Ciphertext, vaultKey) {
		t.Fatal("WrapKey() ciphertext matched plaintext")
	}

	plaintext, err := aead.Open(nil, wrapped.Nonce, wrapped.Ciphertext, nil)
	if err != nil {
		t.Fatalf("aead.Open() error = %v", err)
	}
	if !bytes.Equal(plaintext, vaultKey) {
		t.Fatalf("decrypted wrapped key = %x, want %x", plaintext, vaultKey)
	}
}

func TestWrapKeyRejectsInvalidKEK(t *testing.T) {
	vaultKey := bytes.Repeat([]byte{2}, vaultKeySize)

	if _, err := WrapKey([]byte("short"), vaultKey); err == nil {
		t.Fatal("WrapKey() error = nil, want error")
	}
}
