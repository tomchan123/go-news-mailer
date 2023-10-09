package mailer

import (
	"embed"
	"fmt"
	ht "html/template"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	tt "text/template"
	"time"

	"github.com/tomchan123/go-news-mailer/internal/db"
	"github.com/wneessen/go-mail"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MailServer struct {
	userName   string
	password   string
	host       string
	db         *db.Mongodb
	newsServer *NewsServer
}

const (
	senderAddr     = "noreply@tmc.com"
	senderName     = "TMC Mailer"
	subject        = "AI Newsletter"
	fetchPeriod    = 24
	timeFormat     = "15:04:05"
	timeFullFormat = "2006/01/02 15:04:05"
)

var period = time.Date(0, 0, 0, 23, 16, 0, 0, time.UTC)

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

func now() string {
	return time.Now().Format(timeFormat)
}

func (ms *MailServer) Start() {
	fmt.Printf("<mailer> %s Starting Mail Server\n", now())

	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGTERM, syscall.SIGINT)
		fmt.Printf("<mailer> %s Exiting, Reason: %v\n", now(), <-sigc)
	}()

	for {
		d := ms.getNextScheduleTime()
		fmt.Printf("<mailer> %s Next scehduled email bulk sending on (%s)\n", now(), d.Format(timeFullFormat))
		waitUntil(d)

		go func() {
			fmt.Printf("<mailer> %s Email sending process scheduled on (%s)\n", now(), d.Format(timeFullFormat))

			subs, err := ms.db.GetAllSubscriptions()
			if err != nil {
				fmt.Printf("<mailer> %s Failed to get subscriptions: %v\n", now(), err)
				return
			}

			nas, err := ms.newsServer.FetchRecentArticles(3, fetchPeriod)
			if err != nil {
				fmt.Printf("<mailer> %s Failed to get subscriptions: %v\n", now(), err)
			}

			t := time.Now()
			if err := ms.SendBulkNews(subs, nas); err != nil {
				fmt.Printf("<mailer> %s Failed to send email: %v\n", now(), err)
				return
			}

			fmt.Printf("<mailer> %s Successfully Sent Emails (Count: %d)\n", now(), len(subs))

			var sids []primitive.ObjectID
			for _, sub := range subs {
				sids = append(sids, sub.Uid)
			}
			var vnas []db.NewsAritcle
			for _, na := range nas {
				vnas = append(vnas, *na)
			}
			ms.db.CreateOneEmail(&db.Email{
				SentAt:           t,
				Articles:         vnas,
				SubscriptionUIDs: sids,
			})
		}()
	}
}

func (ms *MailServer) getNextScheduleTime() time.Time {
	now := time.Now()
	nxt := time.Date(now.Year(), now.Month(),
		now.Day(), period.Hour(), period.Minute(),
		period.Second(), period.Nanosecond(), time.UTC)
	if nxt.UTC().Compare(now) <= 0 {
		nxt = nxt.AddDate(0, 0, 1)
	}
	return nxt
}

func waitUntil(d time.Time) {
	t := time.NewTimer(time.Until(d))
	defer func() {
		if !t.Stop() {
			<-t.C
		}
	}()

	<-t.C
}

func (ms *MailServer) SendBulkNews(ss []*db.Subscription, nas []*db.NewsAritcle) error {
	var mes []*mail.Msg
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, s := range ss {
		rn := r.Int31()
		m, err := genNewsEmail(s, nas)
		if err != nil {
			return fmt.Errorf("mailer.MailServer.SendBulkNews > failed to gen email: %v", err)
		}
		m.SetBulk()
		if err := m.EnvelopeFrom(fmt.Sprintf("noreply+%d@tmc.com", rn)); err != nil {
			return fmt.Errorf("mailer.MailServer.SendBulkNews > failed to set envelope from: %v", err)
		}

		mes = append(mes, m)
	}

	c, err := mail.NewClient(ms.host,
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithTLSPolicy(mail.TLSMandatory),
		mail.WithUsername(ms.userName),
		mail.WithPassword(ms.password),
	)
	if err != nil {
		return fmt.Errorf("mailer.MailServer.SendBulkNews > failed to create client: %v", err)
	}
	if err := c.DialAndSend(mes...); err != nil {
		return fmt.Errorf("mailer.MailServer.SendBulkNews > failed to send bulk emails: %v", err)
	}

	return nil
}

func (ms *MailServer) SendNews(s *db.Subscription) error {
	// nas, err := ms.newsServer.FetchRecentArticles(3, period)
	// if err != nil {
	// 	return fmt.Errorf("mailer.SendNews > failed to fetch news: %v", err)
	// }

	// dummy
	nas := []*db.NewsAritcle{
		{
			AISummary:   "Lorem ipsum dolor sit amet consectetur adipisicing elit. Suscipit libero magnam nulla accusamus, animi accusantium. Illo vel minus corporis totam. Ipsum nostrum cupiditate dicta reprehenderit assumenda, saepe amet non iure!",
			Title:       "The Greatest Title Ever",
			Url:         "https://google.com",
			PublishedAt: time.Now(),
			ImgUrl:      "https://imagizer.imageshack.com/img540/6610/a60fa8.jpg",
		},
	}
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

func genNewsEmail(s *db.Subscription, news []*db.NewsAritcle) (*mail.Msg, error) {
	_, err := tt.New("text_template").Parse(newsTextTemplate)
	if err != nil {
		return nil, fmt.Errorf("mailer.getNewsEmailMsg > failed to set text template: %v", err)
	}
	htmpl, err := ht.New("html_template").Parse(newsHtmlTemplate)
	if err != nil {
		return nil, fmt.Errorf("mailer.getNewsEmailMsg > failed to set html template: %v", err)
	}

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
		Ns []*db.NewsAritcle
		T  time.Time
	}{
		s,
		news,
		time.Now(),
	}
	if err := m.EmbedFromEmbedFS("assets/logo.png", &logo); err != nil {
		return nil, fmt.Errorf("mailer.getNewsEmailMsg > failed to embed logo file: %v", err)
	}
	if err := m.SetBodyHTMLTemplate(htmpl, v); err != nil {
		return nil, fmt.Errorf("mailer.getNewsEmailMsg > failed to set html template: %v", err)
	}
	// if err := m.AddAlternativeTextTemplate(ttmpl, v); err != nil {
	// 	return nil, fmt.Errorf("mailer.getNewsEmailMsg > failed to set text template: %v", err)
	// }

	return m, nil
}
