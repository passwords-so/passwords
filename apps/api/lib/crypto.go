package lib

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
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

func GenerateKeyPair() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}

	publicKey := &privateKey.PublicKey

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}

	privateKeyB64 := base64.StdEncoding.EncodeToString(privateKeyBytes)
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKeyBytes)

	return privateKeyB64, publicKeyB64, nil
}
