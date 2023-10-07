package news

import (
	"context"
	"fmt"
	"log"

	"github.com/tomchan123/go-news-mailer/internal/db"
	"github.com/tomchan123/go-news-mailer/internal/gen/email"
	"github.com/tomchan123/go-news-mailer/internal/mailer"
)

type emailsrvc struct {
	logger     *log.Logger
	mailServer *mailer.MailServer
}

// NewEmail returns the email service implementation.
func NewEmail(logger *log.Logger, ms *mailer.MailServer) email.Service {
	return &emailsrvc{logger, ms}
}

func (e *emailsrvc) SendOneEmail(ctx context.Context, p *email.SubscriptionSendOneEmailPayload) (err error) {
	e.logger.Print("email.SendOneEmail")

	s := db.Subscription{
		Name: "Sir/Madam",
	}
	if p.Name != nil {
		s.Name = *p.Name
	}
	s.Email = p.Email

	if e := e.mailServer.SendNews(&s); e != nil {
		err = email.MakeServerError(fmt.Errorf("server error: %v", e))
		return
	}

	return
}
