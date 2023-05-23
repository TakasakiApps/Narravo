package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
