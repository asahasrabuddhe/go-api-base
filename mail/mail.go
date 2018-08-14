package mail

import (
	"fmt"
	"strings"
)

type Mail struct {
	FromId  string
	ToIds   []string
	CcIds   []string
	BccIds  []string
	Subject string
	Body    string
}

type SmtpServer struct {
	Host string
	Port string
}

func (s *SmtpServer) ServerName() string {
	return s.Host + ":" + s.Port
}

func (mail *Mail) BuildMessage() (message string) {
	message += fmt.Sprintf("From: %s\r\n", mail.FromId)

	if len(mail.ToIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.ToIds, ";"))
	}

	if len(mail.CcIds) > 0 {
		message += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.CcIds, ";"))
	}

	if len(mail.BccIds) > 0 {
		message += fmt.Sprintf("Bcc: %s\r\n", strings.Join(mail.BccIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	message += "\r\n" + mail.Body

	return
}
