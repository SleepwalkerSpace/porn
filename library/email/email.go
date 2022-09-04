package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

type EmailMgr struct {
	name     string
	smtpHost string
	smtpPort int
	smtpUser string
	smtpPswd string

	pool *email.Pool
	ch   chan *email.Email
}

func New(name, smtpHost string, smtpPort int, smtpUser, smtpPswd string, count int, log interface {
	Errorf(format string, args ...interface{})
}) (*EmailMgr, error) {

	pool, err := email.NewPool(
		fmt.Sprintf("%v:%v", smtpHost, smtpPort),
		count,
		smtp.PlainAuth("", smtpUser, smtpPswd, smtpHost),
		&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         smtpHost,
		},
	)
	if err != nil {
		return nil, err
	}

	emgr := &EmailMgr{
		name:     name,
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		smtpUser: smtpUser,
		smtpPswd: smtpPswd,
		pool:     pool,
		ch:       make(chan *email.Email, 24),
	}

	for i := 0; i < count; i++ {
		go func() {
			for e := range emgr.ch {
				if err := emgr.pool.Send(e, time.Second*5); err != nil {
					log.Errorf("email:%+v error:%v", *e, err)
				}
			}
		}()
	}
	return emgr, nil
}

func (emgr *EmailMgr) Send(subject, text, html, to string) {
	e := email.NewEmail()
	e.From = emgr.smtpUser
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)
	e.HTML = []byte(html)
	emgr.ch <- e
}
