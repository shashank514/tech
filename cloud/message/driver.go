package message

type MessageDriver interface {
	SendEmail(toEmail string, subject, body string) error
	SendEmailWithPDF(pdfData []byte, toEmail string, subject, body string) error
}
