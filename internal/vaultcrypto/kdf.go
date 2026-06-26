package vaultcrypto

// KDFParams records how the master password is turned into key material.
type KDFParams struct {
	Algorithm string
	Salt      []byte
	MemoryKiB uint32
	Time      uint32
	Threads   uint8
	KeyLen    uint32
}

// DeriveKey derives key material from a password using the given KDF parameters.
func DeriveKey(password []byte, params KDFParams) ([]byte, error)
