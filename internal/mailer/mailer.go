package mailer

import (
	"embed"
	"fmt"
	"os"
	tt "text/template"
	"time"

	"github.com/tomchan123/go-news-mailer/internal/db"
	"github.com/wneessen/go-mail"
)

type MailServer struct {
	userName   string
	password   string
	host       string
	db         *db.Mongodb
	newsServer *NewsServer
}

const (
	senderAddr = "noreply@tmc.com"
	senderName = "TMC Mailer"
	subject    = "AI Newsletter"
	period     = 24
)

//go:embed assets/logo.png
var logo embed.FS

func CreateMailer(db *db.Mongodb, ns *NewsServer) *MailServer {
	return &MailServer{
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PWD"),
		"smtp.gmail.com",
		db,
		ns,
	}
}

func (ms *MailServer) Start() {

}

func SendBulkNews(s []*db.Subscription) (int, error) {
	return 0, nil
}

func (ms *MailServer) SendNews(s *db.Subscription) error {
	nas, err := ms.newsServer.FetchRecentArticles(3, period)
	if err != nil {
		return fmt.Errorf("mailer.SendNews > failed to fetch news: %v", err)
	}

	// dummy
	// nas := []*NewsAritcle{
	// 	{
	// 		AISummary:   "testing",
	// 		Title:       "title",
	// 		PublishedAt: time.Now(),
	// 	},
	// }
	m, err := genNewsEmail(s, nas)
	if err != nil {
		return fmt.Errorf("mailer.SendNews > failed to get email msg: %v", err)
	}

	c, err := mail.NewClient(ms.host,
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithTLSPolicy(mail.TLSMandatory),
		mail.WithUsername(ms.userName),
		mail.WithPassword(ms.password),
	)
	if err != nil {
		return fmt.Errorf("mailer.SendNews > failed to get email msg: %v", err)
	}
	if err := c.DialAndSend(m); err != nil {
		return fmt.Errorf("mailer.SendNews > failed to send email: %v", err)
	}

	return nil
}

func genNewsEmail(s *db.Subscription, news []*NewsAritcle) (*mail.Msg, error) {
	ttmpl, err := tt.New("text_template").Parse(newsTextTemplate)
	if err != nil {
		return nil, fmt.Errorf("mailer.getNewsEmailMsg > failed to set text template: %v", err)
	}
	// htmpl, err := ht.New("html_template").Parse(newsHtmlTemplate)
	// if err != nil {
	// 	return nil, fmt.Errorf("mailer.getNewsEmailMsg > failed to set html template: %v", err)
	// }

	m := mail.NewMsg()
	m.SetMessageID()
	m.SetDate()
	if err := m.FromFormat(senderName, senderAddr); err != nil {
		return nil, fmt.Errorf("mailer.getNewsEmailMsg > failed to set from: %v", err)
	}
	if err := m.AddToFormat(s.Name, s.Email); err != nil {
		return nil, fmt.Errorf("mailer.getNewsEmailMsg > failed to add to: %v", err)
	}
	m.Subject(subject)
	v := struct {
		S  *db.Subscription
		Ns []*NewsAritcle
		T  time.Time
	}{
		s,
		news,
		time.Now(),
	}
	if err := m.EmbedFromEmbedFS("assets/logo.png", &logo); err != nil {
		return nil, fmt.Errorf("mailer.getNewsEmailMsg > failed to embed logo file: %v", err)
	}
	// if err := m.SetBodyHTMLTemplate(htmpl, v); err != nil {
	// 	return nil, fmt.Errorf("mailer.getNewsEmailMsg > failed to set html template: %v", err)
	// }
	if err := m.AddAlternativeTextTemplate(ttmpl, v); err != nil {
		return nil, fmt.Errorf("mailer.getNewsEmailMsg > failed to set text template: %v", err)
	}

	return m, nil
}
