package email

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"

	"github.com/jonylim/basego/internal/pkg/common/constant/envvar"
	"github.com/jonylim/basego/internal/pkg/common/helper"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

type config struct {
	Host, Port, Username, Password, FromEmail, FromName string
}

// CRLF is Carriage Return Line Feed.
const CRLF = "\r\n"

var errInit = errors.New("Email sender is not initialized")
var cfg config
var env string

var smtpAddr string
var smtpAuth smtp.Auth

// Init initializes the email sender.
func Init() {
	// Load configs from environment variables.
	cfg = config{
		Host:      os.Getenv(envvar.SMTP.Host),
		Port:      os.Getenv(envvar.SMTP.Port),
		Username:  os.Getenv(envvar.SMTP.Username),
		Password:  os.Getenv(envvar.SMTP.Password),
		FromEmail: os.Getenv(envvar.SMTP.FromEmail),
		FromName:  os.Getenv(envvar.SMTP.FromName),
	}
	logger.Println("email", fmt.Sprintf(`Host = %s, Port = %s, Username = %s, Password = %v, FromEmail = %s, FromName = %s`,
		cfg.Host, cfg.Port, cfg.Username, cfg.Password != "", cfg.FromEmail, cfg.FromName))

	// Validate the configs.
	isValid := true
	if cfg.Host == "" {
		logger.Println("email", "WARN: SMTP host is empty")
		isValid = false
	}
	if cfg.Port == "" {
		cfg.Port = "587"
		logger.Println("email", fmt.Sprintf("WARN: SMTP port is empty, set to '%v' as default", cfg.Port))
	} else if _, err := helper.StringToInt(cfg.Port); err != nil {
		logger.Println("email", fmt.Sprintf("ERROR: SMTP port '%v' is invalid", cfg.Port))
		isValid = false
	}
	if cfg.Username == "" {
		logger.Println("email", "WARN: SMTP username is empty")
		isValid = false
	}
	if cfg.Password == "" {
		logger.Println("email", "WARN: SMTP password is empty")
	}
	if cfg.FromEmail == "" {
		if cfg.Username == "" {
			logger.Println("email", "WARN: SMTP FROM email address is empty")
		} else {
			cfg.FromEmail = cfg.Username
			logger.Println("email", fmt.Sprintf("WARN: SMTP FROM email address is empty, set to '%v' as default", cfg.FromEmail))
		}
	}
	if cfg.FromName == "" {
		if cfg.FromEmail == "" {
			logger.Println("email", "WARN: SMTP FROM name is empty")
		} else {
			cfg.FromName = cfg.FromEmail
			logger.Println("email", fmt.Sprintf("WARN: SMTP FROM name is empty, set to '%v' as default", cfg.FromName))
		}
	}
	if isValid {
		env = os.Getenv(envvar.Environment)

		// Run the worker.
		run()

		errInit = nil
		smtpAddr = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
		smtpAuth = smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
		if env != "production" {
			smtpAuth = &unencryptedAuth{smtpAuth}
		}
	} else {
		errInit = errors.New("Email sender is not configured correctly")
	}
}
