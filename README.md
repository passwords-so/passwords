# passwords

a local only offline password manager

## Security idea

Each vault has a user-provided master password which is turned into a key using a KDF, argon2id. The vault also generates a random vault key, which is wrapped by the master key and stored in the vault. The vault items are encrypted using the vault key. This way, the master password is never stored, and the vault key is only accessible if the correct master password is provided.

User derived key will also be used to encrypt the rest of the data associated with the vault, such as metadata, settings, audit logs, etc. Ideally without the master password the entire DB is just a blob of random data.
