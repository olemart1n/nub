// Package session provies logic for authentication
package session

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSessionID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}
