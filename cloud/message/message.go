package message

import (
	"encoding/base64"
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

func (t *Message) SendEmailWithPDF(pdfData []byte, toEmail string, subject, body string) error {
	// Set up authentication information
	from := "shashankkp304@gmail.com"
	password := "devc tptr xwru imyu"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Set up authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)
	addr := smtpHost + ":" + smtpPort

	// Create the email message with MIME formatting
	message := fmt.Sprintf("From: %s\r\n", from)
	message += fmt.Sprintf("To: %s\r\n", toEmail)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += "MIME-Version: 1.0\r\n"
	message += "Content-Type: multipart/mixed; boundary=boundary42\r\n"
	message += "\r\n--boundary42\r\n"
	message += "Content-Type: text/plain; charset=utf-8\r\n"
	message += "\r\n" + body + "\r\n"
	message += "\r\n--boundary42\r\n"
	message += "Content-Type: application/pdf\r\n"
	message += "Content-Disposition: attachment; filename=\"investment_holding.pdf\"\r\n"
	message += "Content-Transfer-Encoding: base64\r\n"
	message += "\r\n"

	// Encode PDF to base64
	encodedPDF := make([]byte, base64.StdEncoding.EncodedLen(len(pdfData)))
	base64.StdEncoding.Encode(encodedPDF, pdfData)

	message += string(encodedPDF)
	message += "\r\n--boundary42--\r\n"

	err := smtp.SendMail(addr, auth, from, []string{toEmail}, []byte(message))
	if err != nil {
		return err
	}

	fmt.Println("PDF sent successfully!")
	return nil
}
