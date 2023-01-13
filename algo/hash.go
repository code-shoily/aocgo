package algo

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMD5Hash computes the hex encoded hash for `text`
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
