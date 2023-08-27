package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSHA256(text string) string {
	// Create a new SHA-256 hasher
	hasher := sha256.New()

	// Write the input text to the hasher
	hasher.Write([]byte(text))

	// Calculate the hash
	hashedBytes := hasher.Sum(nil)

	// Convert the hash to a hexadecimal string
	hashedString := hex.EncodeToString(hashedBytes)

	return hashedString
}
