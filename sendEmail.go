package main

import (
	"fmt"
	"net/smtp"
)

// Function to send the OTP via email using smtp.SendMail
func SendEmail(toEmail string, otp string, message []byte) error {
	// Set up authentication information
	from := "shashankkp304@gmail.com"
	password := "devc tptr xwru imyu"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Compose the message
	// subject := "Subject: Your OTP Code\r\n"
	// body := fmt.Sprintf("Your OTP code is: %s", otp)
	// msg := []byte(subject + "\r\n" + body + "\r\n")

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
