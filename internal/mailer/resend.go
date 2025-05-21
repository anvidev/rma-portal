package mailer

import (
	"bytes"
	"html/template"

	"github.com/resend/resend-go/v2"
)

type resendMailer struct {
	client *resend.Client
	from   string // format: "name <email>"
}

func NewResendMailer(apikey, from string) Mailer {
	return &resendMailer{
		from:   from,
		client: resend.NewClient(apikey),
	}
}

func (m *resendMailer) Send(to []string, email string, data any) error {
	tmpl, err := template.ParseFS(templates, email)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(subject, "subject", data); err != nil {
		return err
	}

	body := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(body, "body", data); err != nil {
		return err
	}

	mail := &resend.SendEmailRequest{
		From:    m.from,
		To:      to,
		Html:    body.String(),
		Subject: subject.String(),
	}

	if _, err = m.client.Emails.Send(mail); err != nil {
		return err
	}

	return nil
}
