package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// SHA256inHex generates SHA-256 hash from str and returns it as hex string.
func SHA256inHex(str string) string {
	b := sha256.Sum256([]byte(str))
	return hex.EncodeToString(b[:])
}

// SHA512inHex generates SHA-512 hash from str and returns it as hex string.
func SHA512inHex(str string) string {
	b := sha512.Sum512([]byte(str))
	return hex.EncodeToString(b[:])
}

// MD5inHex generates MD5 hash from str and returns it as hex string.
func MD5inHex(str string) string {
	b := md5.Sum([]byte(str))
	return hex.EncodeToString(b[:])
}
