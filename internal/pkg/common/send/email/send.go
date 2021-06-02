package email

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// Recipients defines a message's recipients.
type Recipients struct {
	To, Cc, Bcc []string
}

type sendJob struct {
	Attempt    int
	Message    Message
	Recipients Recipients
}

const (
	numWorker  = 3
	maxAttempt = 3
)

var ch chan *sendJob

func run() {
	if ch != nil {
		logger.Error("email", "Already running")
		return
	}
	logger.Println("email", fmt.Sprintf("Running %d workers...", numWorker))
	ch = make(chan *sendJob, numWorker)
	for i := 0; i < numWorker; i++ {
		go func() {
			for {
				select {
				case j := <-ch:
					logger.Println("email", fmt.Sprintf("doSend: Attempt = %d, Subject = %s, To = %v, Cc = %v, Bcc = %v",
						j.Attempt, j.Message.subject, j.Recipients.To, j.Recipients.Cc, j.Recipients.Bcc))
					if err := doSend(j.Message, j.Recipients); err != nil {
						// Requeue the job.
						if j.Attempt < maxAttempt {
							go nextAttempt(j)
						}
					}
				}
			}
		}()
	}
}

// Send sends an email to the specified recipients.
func Send(msg Message, recipients Recipients) error {
	if errInit != nil {
		return errInit
	}
	if len(recipients.To) == 0 {
		return errors.New("email: recipients can't be empty")
	}
	ch <- &sendJob{1, msg, recipients}
	return nil
}

func nextAttempt(j *sendJob) {
	time.Sleep(time.Second)
	ch <- &sendJob{j.Attempt + 1, j.Message, j.Recipients}
}

func doSend(msg Message, recipients Recipients) error {
	subject := msg.subject
	if env != "production" {
		subject = fmt.Sprintf(`*%s* %s`, env, subject)
	}
	var hdrContentType string
	if msg.contentType != "" {
		hdrContentType = fmt.Sprintf(`Content-Type: %s`, msg.contentType) + CRLF
	}
	var cc string
	if len(recipients.Cc) != 0 {
		cc = fmt.Sprintf(`Cc: %s`, strings.Join(recipients.Cc, ",")) + CRLF
	}
	content := fmt.Sprintf(`From: %s <%s>`, cfg.FromName, cfg.FromEmail) + CRLF +
		fmt.Sprintf(`To: %s`, strings.Join(recipients.To, ",")) + CRLF +
		cc +
		fmt.Sprintf(`Subject: %s`, subject) + CRLF +
		hdrContentType +
		CRLF +
		msg.body
	fmt.Println(content)
	to := make([]string, 0, len(recipients.To)+len(recipients.Cc)+len(recipients.Bcc))
	to = append(append(append(to, recipients.To...), recipients.Cc...), recipients.Bcc...)
	err := smtp.SendMail(smtpAddr, smtpAuth, cfg.FromEmail, to, []byte(content))
	if err != nil {
		logger.Error("email", logger.FromError(err))
	}
	return err
}
