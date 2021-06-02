package email

import (
	"fmt"
	"net/smtp"

	"github.com/jonylim/basego/internal/pkg/common/logger"
)

type unencryptedAuth struct {
	smtp.Auth
}

func (a *unencryptedAuth) Start(server *smtp.ServerInfo) (proto string, toServer []byte, err error) {
	if server.TLS {
		return a.Auth.Start(server)
	}
	logger.Trace("email", fmt.Sprintf("Ignore unencrypted connection to %s", server.Name))
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}
