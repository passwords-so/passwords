package utils

import "github.com/google/uuid"

// GenerateID generates a UUIDv4 with an optional prefix.
func GenerateID(prefix ...string) string {
	id := uuid.New().String()
	if len(prefix) > 0 {
		id = prefix[0] + "_" + id
	}
	return id
}
