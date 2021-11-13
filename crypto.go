package common

import "crypto/rand"

// RandomBytes generates a slice of [length] random bytes.
func RandomBytes(length int) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)

	return key, err
}
