package clientapi

import (
	"encoding/base64"
	"encoding/json"
)

// -----------------------------------------------------------------------------
// Email Verification Token
// -----------------------------------------------------------------------------

type emailVerificationToken struct {
	ID    int64  `json:"otpID"`
	Key   string `json:"otpKey"`
	Code  string `json:"otpCode"`
	Email string `json:"email"`
}

func (t *emailVerificationToken) Encode() (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func (t *emailVerificationToken) Decode(s string) error {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, t)
}

// -----------------------------------------------------------------------------
// Reset Password Token
// -----------------------------------------------------------------------------

type emailResetPasswordToken struct {
	ID    int64  `json:"otpID"`
	Key   string `json:"otpKey"`
	Code  string `json:"otpCode"`
	Email string `json:"email"`
}

func (t *emailResetPasswordToken) Encode() (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func (t *emailResetPasswordToken) Decode(s string) error {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, t)
}
