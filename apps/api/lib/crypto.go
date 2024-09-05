package lib

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// Combine salt and hash
	encodedHash := base64.RawStdEncoding.EncodeToString(append(salt, hash...))

	return encodedHash, nil
}

func ComparePasswordAndHash(password, encodedHash string) (bool, error) {
	decoded, err := base64.RawStdEncoding.DecodeString(encodedHash)
	if err != nil {
		return false, err
	}

	salt, hash := decoded[:16], decoded[16:]

	newHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	return string(hash) == string(newHash), nil
}
