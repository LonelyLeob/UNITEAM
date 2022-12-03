package api

import (
	"net/smtp"
)

type MailClient struct {
	account  string
	password string
	host     string
	port     string
	to       []string
}

func NewMailClient(account, password string) *MailClient {
	return &MailClient{
		account:  account,
		password: password,
		host:     "smtp.gmail.com",
		port:     "587",
	}
}

func (m *MailClient) SendMail(theme string, body string, mails ...string) error {

	m.to = mails
	address := m.host + ":" + m.port

	message := []byte(theme + body)

	auth := smtp.PlainAuth("", m.account, m.password, m.host)

	if err := smtp.SendMail(address, auth, m.account, mails, message); err != nil {
		return err
	}

	return nil
}
