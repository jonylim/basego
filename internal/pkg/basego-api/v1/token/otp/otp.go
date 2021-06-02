package otp

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/jonylim/basego/internal/pkg/common/crypto/hash"
)

// Defines OTP methods.
const (
	MethodEmail = "email"
	MethodPhone = "phone"
)

// Defines OTP actions.
const (
	ActionLogin         = "login"
	ActionVerifyEmail   = "verifyEmail"
	ActionVerifyPhone   = "verifyPhone"
	ActionResetPassword = "resetPassword"
)

// TTL defines TTL duration of an OTP, in seconds.
const TTL = 600

// Length defines OTP code's length.
const Length = 6

const (
	charsNumeric      = "0123456789"
	charsAlphanumeric = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var incrementSeed int64

// GenerateAlphanumeric returns a new OTP key and code. The code is a alphanumeric string.
func GenerateAlphanumeric() (otpKey, otpCode string) {
	otpKey, otpCode = generate(charsAlphanumeric)
	return
}

// GenerateNumeric returns a new OTP key and code. The code is a numeric string.
func GenerateNumeric() (otpKey, otpCode string) {
	otpKey, otpCode = generate(charsNumeric)
	return
}

func generate(chars string) (otpKey, otpCode string) {
	incrementSeed++
	if incrementSeed > 9 {
		incrementSeed = 0
	}
	nowNano := time.Now().UnixNano()
	rand.Seed((nowNano/10)*10 + incrementSeed)
	max := len(chars)
	var sb strings.Builder
	for i := 0; i < Length; i++ {
		c := rand.Intn(max)
		sb.WriteString(string(chars[c]))
	}
	otpCode = sb.String()
	otpKey = hash.MD5inHex(fmt.Sprintf("%v+%v=%s", nowNano, incrementSeed, otpCode))
	return
}
