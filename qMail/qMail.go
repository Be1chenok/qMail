package qmail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type Sendner interface {
	Send(smtpAddres SmtpAddres, msg Message, rcvr Receivers) error
}

type sendner struct {
	name     string
	addres   string
	password string
}

func New(name, addres, password string) Sendner {
	return &sendner{
		name:     name,
		addres:   addres,
		password: password,
	}
}

func NewMessage(subject, content string, attachFiles []string) Message {
	return Message{
		Subject:     subject,
		Content:     content,
		AttachFiles: attachFiles,
	}
}

func NewReceivers(to, cc, bcc []string) Receivers {
	return Receivers{
		To:  to,
		Cc:  cc,
		Bcc: bcc,
	}
}

func NewSMTP(domain, port string) SmtpAddres {
	return SmtpAddres{
		Domain: domain,
		Port:   port,
	}
}

func (s sendner) Send(smtpAddres SmtpAddres, msg Message, rcvr Receivers) error {
	email := email.NewEmail()
	email.From = fmt.Sprintf("%s <%s>", s.name, s.addres)
	email.Subject = msg.Subject
	email.HTML = []byte(msg.Content)
	email.To = rcvr.To
	email.Cc = rcvr.Cc
	email.Bcc = rcvr.Bcc

	for _, file := range msg.AttachFiles {
		if _, err := email.AttachFile(file); err != nil {
			return fmt.Errorf("failed to attach file %s: %w", file, err)
		}
	}
	smtpAuth := smtp.PlainAuth("", s.name, s.password, smtpAddres.Domain)

	if err := email.Send(smtpAddres.Domain+":"+smtpAddres.Port, smtpAuth); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
