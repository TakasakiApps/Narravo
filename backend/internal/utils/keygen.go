package utils

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/ohanakogo/exceptiongo"
)

// GenerateAESKey generates and returns a new 96-bit AES key encoded in base64 format.
func GenerateAESKey() string {
	key := make([]byte, 12)
	if _, err := rand.Read(key); err != nil {
		exceptiongo.QuickThrow[types.RuntimeException](err)
	}
	return base64.StdEncoding.EncodeToString(key)
}
