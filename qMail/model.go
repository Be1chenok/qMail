package qmail

type SmtpAddres struct {
	Domain string
	Port   int
}

type Message struct {
	Subject     string
	Content     string
	AttachFiles []string
}

type Receivers struct {
	To  []string
	Cc  []string
	Bcc []string
}
