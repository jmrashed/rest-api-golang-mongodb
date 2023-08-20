package util

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateRandomID generates a random hexadecimal ID with the given length.
func GenerateRandomID(length int) (string, error) {
	bytes := make([]byte, length/2)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
