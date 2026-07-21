package vaultcrypto

import "crypto/rand"

const (
	passwordSlotVersion = 1
	vaultKeySize        = 32
	passwordSlotAAD     = "passwords:v1:vault-key:"
)

// CreatePasswordSlot creates a random vault key protected by a master password.
func CreatePasswordSlot(password []byte, vaultID string) (vaultKey []byte, slot PasswordSlot, err error) {
	// Create new Argon2id parameters with a random salt.
	// Derive a temporary key-encryption key from the master password.
	// Generate a random 32-byte vault key.
	// Seal the vault key and bind it to the vault ID with associated data.
	// Clear the temporary key-encryption key before returning.

	params, err := newKDFParams()
	if err != nil {
		return nil, PasswordSlot{}, err
	}

	tmpKey, err := deriveKey(password, params)
	if err != nil {
		return nil, PasswordSlot{}, err
	}
	defer func() {
		clear(tmpKey)
	}()

	vaultKey = make([]byte, vaultKeySize)
	if _, err := rand.Read(vaultKey); err != nil {
		return nil, PasswordSlot{}, err
	}

	wrappedVaultKey, err := seal(
		tmpKey,
		vaultKey,
		[]byte(passwordSlotAAD+vaultID),
	)
	if err != nil {
		clear(vaultKey)
		return nil, PasswordSlot{}, err
	}

	return vaultKey, PasswordSlot{
		Version:         passwordSlotVersion,
		KDF:             params,
		WrappedVaultKey: wrappedVaultKey,
	}, nil
}

// OpenPasswordSlot unlocks a stored vault key with a master password.
func OpenPasswordSlot(password []byte, vaultID string, slot PasswordSlot) ([]byte, error)

// Derive the key-encryption key using the slot's stored KDF parameters.
// Open the wrapped vault key using the vault ID as associated data.
// Clear the temporary key-encryption key before returning.
// Treat a wrong password and damaged ciphertext as the same unlock failure.
