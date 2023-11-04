package qmail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type QMail interface {
	NewSendner(name, addres, password string) sendner
	NewMessage(subject, content string, attachFiles []string) message
	NewReceivers(to, cc, bcc []string) receivers
	Send(sndr sendner, msg message, rcvr receivers) error
}

type qMail struct {
	smtpDomain string
	smtpPort   string
}

func New(smtpDomain, smtpPort string) QMail {
	return &qMail{
		smtpDomain: smtpDomain,
		smtpPort:   smtpPort,
	}
}

func (q *qMail) NewSendner(name, addres, password string) sendner {
	return sendner{
		Name:     name,
		Addres:   addres,
		Password: password,
	}
}

func (q *qMail) NewMessage(subject, content string, attachFiles []string) message {
	return message{
		Subject:     subject,
		Content:     content,
		AttachFiles: attachFiles,
	}
}

func (q *qMail) NewReceivers(to, cc, bcc []string) receivers {
	return receivers{
		To:  to,
		Cc:  cc,
		Bcc: bcc,
	}
}

func (q qMail) Send(sndr sendner, msg message, rcvr receivers) error {
	email := email.NewEmail()
	email.From = fmt.Sprintf("%s <%s>", sndr.Name, sndr.Addres)
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

	smtpAuth := smtp.PlainAuth("", sndr.Addres, sndr.Password, q.smtpDomain)

	if err := email.Send(q.smtpDomain+":"+q.smtpPort, smtpAuth); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
