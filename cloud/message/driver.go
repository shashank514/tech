package message

type MessageDriver interface {
	SendEmail(toEmail string, subject, body string) error
}
