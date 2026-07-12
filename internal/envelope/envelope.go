package envelope

// Envelope is an encrypted byte payload plus the metadata needed to decrypt it.
type Envelope struct {
	Algorithm  string
	Nonce      []byte
	Ciphertext []byte
}

// Encrypt seals plaintext with authenticated encryption.
func Encrypt(key []byte, plaintext []byte, aad []byte) (Envelope, error)

// Decrypt opens an encrypted envelope and verifies its authenticated data.
func Decrypt(key []byte, envelope Envelope, aad []byte) ([]byte, error)
