package password

import (
	"math/rand"
	"strings"
	"time"

	"github.com/jonylim/basego/internal/pkg/common/crypto/hash"
)

const saltChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz.~!@#$^&*-_=+:;"

// SaltLength defines the length of a salt string.
var SaltLength = 10

var incrementSeed int64

func generateSalt(length int) string {
	incrementSeed++
	if incrementSeed > 9 {
		incrementSeed = 0
	}
	max := len(saltChars)
	rand.Seed((time.Now().UnixNano()/10)*10 + incrementSeed)
	var sb strings.Builder
	for i := 0; i < length; i++ {
		c := rand.Intn(max)
		sb.WriteString(string(saltChars[c]))
	}
	return sb.String()
}

// GenerateSalt generates a random salt string for hashing password.
func GenerateSalt() string {
	return generateSalt(SaltLength)
}

// HashWithSalt hashes a plain text password with salt using SHA-512.
func HashWithSalt(plain, salt string) string {
	return hash.SHA512inHex(salt + plain)
}

// Hash hashes a plain text password with randomly generated salt using SHA-512.
func Hash(plain string) (hashed, salt string) {
	salt = generateSalt(SaltLength)
	hashed = hash.SHA512inHex(salt + plain)
	return
}
