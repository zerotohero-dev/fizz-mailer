package postman

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/pkg/errors"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	// "github.com/zerotohero-dev/fizz-mailer/internal/service"
	"github.com/zerotohero-dev/fizz-mailer/internal/service/postman/template"
	"time"
)

// TODO: to env
const from = "Volkan Özçelik <volkan@hermes.fizzbuzz.pro>"

type Args struct {
	MailgunDomain string
	MailgunApiKey string
	IsDevelopment bool
	Email         string
	Name          string
}

func RelayEmailVerificationMessage(
	args Args, emailVerificationBaseUrl, emailVerificationToken string,
) error {
	body := template.EmailVerificationMessageBody(template.EmailVerificationMessageParams{
		Email:                    args.Email,
		Name:                     args.Name,
		EmailVerificationBaseUrl: emailVerificationBaseUrl,
		Token:                    emailVerificationToken,
	})

	domain := args.MailgunDomain
	apiKey := args.MailgunApiKey

	subject := fmt.Sprintf("[FizzBuzz Pro] %s, please verify your email 🐢", args.Name)

	if args.IsDevelopment {
		log.Info("mailer: %s", subject)
		log.Info("mailer: %s", args.Email)
		log.Info("mailer: %s", body)
		return nil
	}

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, args.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"RelayEmailVerificationMessage: problem sending activation email (%s)",
				log.RedactEmail(args.Email),
			),
		)
	}

	return nil
}

func RelayWelcomeMessage(args Args) error {
	body := template.WelcomeMessageBody(template.WelcomeMessageParams{
		Name: args.Name,
	})

	domain := args.MailgunDomain
	apiKey := args.MailgunApiKey

	subject := fmt.Sprintf("[FizzBuzz Pro] %s, Welcome to %s 🐢", args.Name, "the Jungle")

	if args.IsDevelopment {
		log.Info("mailer: %s", subject)
		log.Info("mailer: %s", args.Email)
		log.Info("mailer: %s", body)
		return nil
	}

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, args.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"RelayWelcomeMessage: problem sending welcome email (%s)",
				log.RedactEmail(args.Email),
			),
		)
	}

	return nil
}

func RelayPasswordResetMessage(args Args, passwordResetBaseUrl, passwordResetToken string) error {
	body := template.PasswordResetMessageBody(template.PasswordResetMessageParams{
		Name:                 args.Name,
		Email:                args.Email,
		Token:                passwordResetToken,
		PasswordResetBaseUrl: passwordResetBaseUrl,
	})

	domain := args.MailgunDomain
	apiKey := args.MailgunApiKey

	subject := fmt.Sprintf("[FizzBuzz Pro] Hi %s, here are your password reset instructions.", args.Name)

	if args.IsDevelopment {
		log.Info("mailer: %s", subject)
		log.Info("mailer: %s", args.Email)
		log.Info("mailer: %s", body)
		return nil
	}

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, args.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"RelayPasswordResetMessage: problem sending password reset email (%s)",
				log.RedactEmail(args.Email),
			),
		)
	}

	return nil
}

func RelayPasswordResetConfirmationMessage(args Args) error {
	body := template.PasswordResetConfirmationMessageBody(template.PasswordResetConfirmationMessageParams{
		Name: args.Name,
	})

	domain := args.MailgunDomain
	apiKey := args.MailgunApiKey

	subject := fmt.Sprintf("[FizzBuzz Pro] Hi %s, you've successfully reset your password.", args.Name)

	if args.IsDevelopment {
		log.Info("mailer: %s", subject)
		log.Info("mailer: %s", args.Email)
		log.Info("mailer: %s", body)
		return nil
	}

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, args.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"RelayPasswordResetConfirmationMessage: problem sending passwrod reset confirmation email (%s)",
				log.RedactEmail(args.Email),
			),
		)
	}

	return nil
}

func RelaySubscribedMessage(args Args) error {
	body := template.SubscribedMessageBody(template.SubscribedMessageParams{
		Name: args.Name,
	})

	domain := args.MailgunDomain
	apiKey := args.MailgunApiKey

	subject := fmt.Sprintf("[FizzBuzz Pro] Hi %s, welcome to the jungle.", args.Name)

	if args.IsDevelopment {
		log.Info("mailer: %s", subject)
		log.Info("mailer: %s", args.Email)
		log.Info("mailer: %s", body)
		return nil
	}

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, args.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"RelaySubscribedMessage: problem sending passwrod reset confirmation email (%s)",
				log.RedactEmail(args.Email),
			),
		)
	}

	return nil
}
