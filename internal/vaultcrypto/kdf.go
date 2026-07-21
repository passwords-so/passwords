package vaultcrypto

import (
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/argon2"
)

var ErrInvalidKDFParams = errors.New("invalid KDF parameters")

const (
	defaultAlgorithm = "argon2id"
	defaultSaltLen   = 16
	defaultMemoryKiB = 64 * 1024
	defaultTime      = 3
	defaultThreads   = 4
	defaultKeyLen    = 32
)

// newKDFParams creates the settings for a new password slot.
func newKDFParams() (KDFParams, error) {
	salt := make([]byte, defaultSaltLen)
	if _, err := rand.Read(salt); err != nil {
		return KDFParams{}, err
	}
	return KDFParams{
		Algorithm: "argon2id",
		Salt:      salt,
		MemoryKiB: defaultMemoryKiB,
		Time:      defaultTime,
		Threads:   defaultThreads,
		KeyLen:    defaultKeyLen,
	}, nil
}

// deriveKey turns a master password into a key-encryption key.
func deriveKey(password []byte, params KDFParams) ([]byte, error) {
	if err := validateKDFParams(params); err != nil {
		return nil, err
	}

	key := argon2.IDKey(
		password,
		params.Salt,
		params.Time,
		params.MemoryKiB,
		params.Threads,
		params.KeyLen,
	)

	return key, nil
}

func validateKDFParams(params KDFParams) error {
	if params.Algorithm != defaultAlgorithm {
		return ErrInvalidKDFParams
	}
	if len(params.Salt) != defaultSaltLen {
		return ErrInvalidKDFParams
	}
	if params.MemoryKiB != defaultMemoryKiB {
		return ErrInvalidKDFParams
	}
	if params.Time != defaultTime {
		return ErrInvalidKDFParams
	}
	if params.Threads != defaultThreads {
		return ErrInvalidKDFParams
	}
	if params.KeyLen != defaultKeyLen {
		return ErrInvalidKDFParams
	}
	return nil
}
