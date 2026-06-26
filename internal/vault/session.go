package vault

// Session is the unlocked, in-memory vault state.
// It exists only after Unlock succeeds and should be cleared by Lock.
type Session struct {
	vault    *Vault
	vaultKey []byte
}
