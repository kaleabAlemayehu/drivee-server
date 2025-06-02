package utils

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"
	"os"

	"github.com/kaleabAlemayehu/drivee-server/template"
)

func SendEmail(email string, userName string, url string) error {
	host := os.Getenv("SMTP_HOST")
	from := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	var buf bytes.Buffer
	template.ResetPasswordEmail(userName, url).Render(context.Background(), &buf)

	htmlBody := buf.String()

	message := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: Reset Your Password\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
			"\r\n%s",
		from, email, htmlBody,
	))
	return smtp.SendMail(fmt.Sprintf("%s:587", host), smtp.PlainAuth("", from, password, host), from, []string{email}, message)
}
