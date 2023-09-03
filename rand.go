package common

import (
	"crypto/rand"
	mRand "math/rand"
	"time"
)

// RandomBytes generates a slice of [length] random bytes.
func RandomBytes(n int) ([]byte, error) {
	key := make([]byte, n)
	_, err := rand.Read(key)

	return key, err
}

// RandomString generates a random string of length n.
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		mRand.New(mRand.NewSource(time.Now().UnixNano()))
		s[i] = letters[mRand.Intn(len(letters))]
	}

	return string(s)
}
