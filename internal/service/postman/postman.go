package postman

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/pkg/errors"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-mailer/internal/service"
	"github.com/zerotohero-dev/fizz-mailer/internal/service/postman/template"
	"time"
)

const from = "Volkan Özçelik <volkan@hermes.fizzbuzz.pro>"

func RelayEmailVerificationMessage(
	args service.Args,
	email, name, emailVerificationToken string,
) error {
	body := template.EmailVerificationMessageBody(template.EmailVerificationMessageParams{
		Email:                    email,
		Name:                     name,
		EmailVerificationBaseUrl: args.EmailVerificationBaseUrl,
		Token:                    emailVerificationToken,
	})

	domain := args.MailgunDomain
	apiKey := args.MailgunApiKey

	subject := fmt.Sprintf("[FizzBuzz Pro] %s, please verify your email 🐢", name)

	if args.IsDevelopment {
		log.Info("mailer: %s", subject)
		log.Info("mailer: %s", email)
		log.Info("mailer: %s", body)
		return nil
	}

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"RelayEmailVerificationMessage: problem sending activation email (%s)",
				log.RedactEmail(email),
			),
		)
	}

	return nil
}

func RelayWelcomeMessage(args service.Args, email, name string) error {
	body := template.WelcomeMessageBody(template.WelcomeMessageParams{
		Name: name,
	})

	domain := args.MailgunDomain
	apiKey := args.MailgunApiKey

	subject := fmt.Sprintf("[FizzBuzz Pro] %s, Welcome to %s 🐢", name)

	if args.IsDevelopment {
		log.Info("mailer: %s", subject)
		log.Info("mailer: %s", email)
		log.Info("mailer: %s", body)
		return nil
	}

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"RelayWelcomeMessage: problem sending welcome email (%s)",
				log.RedactEmail(email),
			),
		)
	}

	return nil
}

func RelayPasswordResetMessage(args service.Args, email, name, passwordResetToken string) error {
	body := template.PasswordResetMessageBody(template.PasswordResetMessageParams{
		Name:                 name,
		Email:                email,
		Token:                passwordResetToken,
		PasswordResetBaseUrl: e.Mailer.PasswordResetBaseUrl,
	})

	domain := e.Mailer.MailgunDomain
	apiKey := e.Mailer.MailgunApiKey

	subject := fmt.Sprintf("[ZeroToHero] Hi %s, here are your password reset instructions.", name)

	if e.Deployment.Type == env.Development {
		log.Info("mailer: %s", subject)
		log.Info("mailer: %s", email)
		log.Info("mailer: %s", body)
		return nil
	}

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"RelayPasswordResetMessage: problem sending password reset email (%s)",
				log.RedactEmail(email),
			),
		)
	}

	return nil
}

func RelayPasswordResetConfirmationMessage(e env.FizzEnv, email, name string) error {
	body := template.PasswordResetConfirmationMessageBody(template.PasswordResetConfirmationMessageParams{
		Name: name,
	})

	domain := e.Mailer.MailgunDomain
	apiKey := e.Mailer.MailgunApiKey

	subject := fmt.Sprintf("[FizzBuzz Pro] Hi %s, you've successfully reset your password.", name)

	if e.Deployment.Type == env.Development {
		log.Info("mailer: %s", subject)
		log.Info("mailer: %s", email)
		log.Info("mailer: %s", body)
		return nil
	}

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"RelayPasswordResetConfirmationMessage: problem sending passwrod reset confirmation email (%s)",
				log.RedactEmail(email),
			),
		)
	}

	return nil
}

func RelaySubscribedMessage(e env.FizzEnv, email, name string) error {
	body := template.SubscribedMessageBody(template.SubscribedMessageParams{
		Name: name,
	})

	domain := e.Mailer.MailgunDomain
	apiKey := e.Mailer.MailgunApiKey

	subject := fmt.Sprintf("[FizzBuzz Pro] Hi %s, welcome to the jungle.", name)

	if e.Deployment.Type == env.Development {
		log.Info("mailer: %s", subject)
		log.Info("mailer: %s", email)
		log.Info("mailer: %s", body)
		return nil
	}

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"RelaySubscribedMessage: problem sending passwrod reset confirmation email (%s)",
				log.RedactEmail(email),
			),
		)
	}

	return nil
}
