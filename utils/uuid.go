package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateHexUUID() string {
	bytes := make([]byte, 14) // 14 bytes = 28 hex characters
	rand.Read(bytes)

	return hex.EncodeToString(bytes)
}
