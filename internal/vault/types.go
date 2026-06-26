package vault

import "time"

// Vault is the decrypted domain model while the vault is unlocked.
type Vault struct {
	ID        string
	Name      string
	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time

	Items map[string]VaultItem
}

// VaultItem is one decrypted password-manager item.
type VaultItem struct {
	ID        string
	Kind      ItemKind
	Title     string
	Username  string
	URL       string
	Password  []byte
	Notes     string
	Tags      []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ItemKind identifies what kind of vault item is stored.
type ItemKind string

const (
	ItemKindLogin ItemKind = "login"
)

// AddLoginInput is the domain input for creating a login item.
type AddLoginInput struct {
	Title    string
	URL      string
	Username string
	Password []byte
	Notes    string
	Tags     []string
}
