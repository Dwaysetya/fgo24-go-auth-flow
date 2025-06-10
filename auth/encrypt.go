package auth

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
