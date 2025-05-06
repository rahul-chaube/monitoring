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
		return fmt.Errorf("template parsing failed: %v", err)
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("template execution failed: %v", err)
	}

	// Hardcoded SMTP credentials (for demo/dev)
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpUser := "your-email@gmail.com"
	smtpPass := "your-app-password"

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n"
	msg += fmt.Sprintf("Subject: %s\n\n%s", subject, body.String())

	// If want to take from env variables:
	// smtpUser = os.Getenv("SMTP_USER")
	// auth = smtp.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"), os.Getenv("SMTP_HOST"))

	return smtp.SendMail(
		fmt.Sprintf("%s:%s", smtpHost, smtpPort),
		auth,
		smtpUser,
		[]string{to},
		[]byte(msg),
	)
}
