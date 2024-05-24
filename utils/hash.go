// utils/hash.go
package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateShortURL(original string) string {
	hash := sha256.Sum256([]byte(original))
	return base64.URLEncoding.EncodeToString(hash[:6])
}
