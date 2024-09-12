package message

import (
	"fmt"
	"net/smtp"
)

type Message struct {
}

func NewMessage() MessageDriver {
	return &Message{}
}

func (t *Message) SendEmail(toEmail string, subject, body string) error {
	// Set up authentication information
	from := "shashankkp304@gmail.com"
	password := "devc tptr xwru imyu"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Compose the message
	message := []byte(subject + "\r\n" + body + "\r\n")

	// Set up authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send the email
	addr := smtpHost + ":" + smtpPort
	err := smtp.SendMail(addr, auth, from, []string{toEmail}, message)
	if err != nil {
		return err
	}

	fmt.Println("OTP sent successfully!")
	return nil
}
