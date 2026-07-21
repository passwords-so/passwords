package vault

// Vault package errors are used by the backend and translated by CLI/UI layers.
var (
	ErrAlreadyUnlocked error
	ErrLocked          error
	ErrNotCreated      error
	ErrWrongPassword   error
	ErrItemNotFound    error
)
