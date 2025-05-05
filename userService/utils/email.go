package utils

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

type EmailTemplateData struct {
	Subject string
	Header  string
	Body    string
}

// SendEmail sends an email to the specified recipient with the given subject and body.
func SendEmail(to, subject, body string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	from := smtpUser
	toList := []string{to}
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, toList, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}

func SendTemplatedEmail(to string, subject string, data EmailTemplateData, templateFile string) error {
	tmpl, err := template.ParseFiles(templateFile)
	//	tmpl, err := template.ParseFiles("templates/forwarding_email.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		return err
	}

	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n"
	msg += fmt.Sprintf("Subject: %s\n\n%s", subject, body.String())

	auth := smtp.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"), os.Getenv("SMTP_HOST"))
	return smtp.SendMail(
		fmt.Sprintf("%s:%s", os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT")),
		auth,
		os.Getenv("SMTP_USER"),
		[]string{to},
		[]byte(msg),
	)
}
