package qmail

type smtpAddres struct {
	Host string
	Port string
}

type message struct {
	Subject     string
	Content     string
	AttachFiles []string
}

type receivers struct {
	To  []string
	Cc  []string
	Bcc []string
}

type sendner struct {
	Name     string
	Addres   string
	Password string
}
