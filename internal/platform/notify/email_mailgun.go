package notify

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/mailgun/mailgun-go/v3"
	"github.com/pkg/errors"
)

type EmailMailgun struct {
	apiKey             string
	domain             string
	senderEmailAddress string
	templateDir        string
}

// NewEmailSmtp creates an implementation of the Email interface used to send email with SMTP.
func NewEmailMailgun(domain string, apiKey string, sharedTemplateDir, senderEmailAddress string) (*EmailMailgun, error) {

	if domain == "" {
		return nil, errors.New("Domain is required.")
	}
	if apiKey == "" {
		return nil, errors.New("API key is required.")
	}
	if senderEmailAddress == "" {
		return nil, errors.New("Sender email address is required.")
	}

	templateDir := filepath.Join(sharedTemplateDir, "emails")
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		return nil, errors.WithMessage(err, "Email template directory does not exist.")
	}

	return &EmailMailgun{
		domain:             domain,
		apiKey:             apiKey,
		templateDir:        templateDir,
		senderEmailAddress: senderEmailAddress,
	}, nil
}

// Verify ensures the provider works.
func (n *EmailMailgun) Verify() error {
	return nil
}

// Send initials the delivery of an email the provided email address.
func (n *EmailMailgun) Send(ctx context.Context, toEmail, subject, templateName string, data map[string]interface{}) error {
	htmlDat, txtDat, err := parseEmailTemplates(n.templateDir, templateName, data)
	if err != nil {
		return errors.Errorf("Error in sending email: %+v", err)
	}

	mg := mailgun.NewMailgun(n.domain, n.apiKey)
	m := mg.NewMessage(
		n.senderEmailAddress,
		subject,
		string(txtDat),
		toEmail,
	)

	m.SetHtml(string(htmlDat))

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	_, _, err = mg.Send(ctx, m)
	return err
}
