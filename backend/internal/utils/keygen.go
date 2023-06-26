package utils

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/ohanakogo/exceptiongo"
	"math/big"
	"strings"
)

// GenerateAESKey generates and returns a new 96-bit AES key encoded in base64 format.
func GenerateAESKey() string {
	return GenerateKey(16)
}

func GenerateKey(len int) string {
	key := make([]byte, len)
	if _, err := rand.Read(key); err != nil {
		exceptiongo.QuickThrow[types.RuntimeException](err)
	}
	return base64.StdEncoding.EncodeToString(key)
}

// GenerateUniqueID generate a unique ID.
func GenerateUniqueID() string {
	id := GenerateStandardID()
	if dao.GetInstance().IsIdRecordExist(id) {
		return GenerateUniqueID()
	}
	return id
}

// GenerateStandardID generates a random identifier with length of 32 characters.
func GenerateStandardID() string {
	return GenerateAnyID(32)
}

// GenerateAnyID generates a random identifier with specified length.
func GenerateAnyID(length int) string {
	// Defines the character set from which the ID will be generated.
	const charset = "0123456789abcdefghijklmnopqrstuvwxyz"
	// Creates a strings.Builder variable to store the generated ID.
	var idBuilder strings.Builder
	// Generates each character of the ID by randomly selecting a character from the charset.
	for i := 0; i < length; i++ {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			// Throw exception if there is an error generating the random number.
			exceptiongo.QuickThrow[types.RuntimeException](err)
		}
		idBuilder.WriteByte(charset[r.Int64()])
	}
	// Returns the generated ID as a string.
	return idBuilder.String()
}
