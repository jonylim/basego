package emailtemplate

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/jonylim/basego/internal/pkg/common/helper"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// VerifyEmailAddress returns template for email "Verify Email Address".
func VerifyEmailAddress(accountName, link, otpCode string, ttlHours int) (subject, body string, err error) {
	var t *template.Template
	t, err = getByFilename("verify-email-address.html")
	if err != nil {
		logger.Fatal("emailtemplate", fmt.Sprintf("VerifyEmailAddress: %v", logger.FromError(err)))
		return
	}
	data := struct{ Title, Name, Code, Link, TTLHours string }{
		Title:    "Please verify your email address",
		Name:     accountName,
		Link:     link,
		Code:     otpCode,
		TTLHours: helper.IntToString(ttlHours),
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, data); err != nil {
		logger.Fatal("emailtemplate", fmt.Sprintf("VerifyEmailAddress: %v", logger.FromError(err)))
		return
	}
	subject, body = data.Title, string(buf.Bytes())
	return
}

// ResetPassword returns template for email "Reset Password Email".
func ResetPassword(accountName, link, otpCode string, ttlHours int) (subject, body string, err error) {
	var t *template.Template
	t, err = getByFilename("reset-password.html")
	if err != nil {
		logger.Fatal("emailtemplate", fmt.Sprintf("ResetPassword: %v", logger.FromError(err)))
		return
	}
	data := struct{ Title, Name, Code, Link, TTLHours string }{
		Title:    "Reset your password",
		Name:     accountName,
		Link:     link,
		Code:     otpCode,
		TTLHours: helper.IntToString(ttlHours),
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, data); err != nil {
		logger.Fatal("emailtemplate", fmt.Sprintf("ResetPassword: %v", logger.FromError(err)))
		return
	}
	subject, body = data.Title, string(buf.Bytes())
	return
}
