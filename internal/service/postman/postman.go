package postman

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/pkg/errors"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-mailer/internal/service/postman/template"
	"time"
)

const from = "Volkan Özçelik <volkan@hermes.fizzbuzz.pro>"

func RelayEmailVerificationMessage(e env.FizzEnv, email, name, emailVerificationToken string) error {
	body := template.EmailVerificationMessageBody(template.EmailVerificationMessageParams{
		Email: email,
		Name:  name,
		Token: emailVerificationToken,
	})

	domain := e.Mailer.MailgunDomain
	apiKey := e.Mailer.MailgunApiKey

	subject := fmt.Sprintf("[FizzBuzz Pro] %s, please verify your email 🦄", name)

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("RelayEmailVerificationMessage: problem sending activation email (%s)", email),
		)
	}

	return nil
}

func RelayEmailVerifiedMessage(e env.FizzEnv, email, name string) error {
	body := template.EmailVerifiedMessageBody(template.EmailVerifiedMessageParams{
		Name: name,
	})

	domain := e.Mailer.MailgunDomain
	apiKey := e.Mailer.MailgunApiKey

	subject := fmt.Sprintf("[FizzBuzz Pro] %s, you have verified your email 🦄", name)

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("RelayEmailVerifiedMessage: problem account verification confirmation email (%s)", email),
		)
	}

	return nil
}

func RelayWelcomeMessage(e env.FizzEnv, email, name string) error {
	body := template.WelcomeMessageBody(template.WelcomeMessageParams{
		Name: name,
	})

	domain := e.Mailer.MailgunDomain
	apiKey := e.Mailer.MailgunApiKey

	subject := fmt.Sprintf("[FizzBuzz Pro] %s, Welcome to %s 🦄", name)

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("RelayWelcomeMessage: problem sending welcome email (%s)", email),
		)
	}

	return nil
}

func RelayPasswordResetMessage(e env.FizzEnv, email, name, passwordResetToken string) error {
	body := template.PasswordResetMessageBody(template.PasswordResetMessageParams{
		Name:  name,
		Email: email,
		Token: passwordResetToken,
	})

	domain := e.Mailer.MailgunDomain
	apiKey := e.Mailer.MailgunApiKey

	subject := fmt.Sprintf("[ZeroToHero] Hi %s, here are your password reset instructions.", name)

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("RelayPasswordResetMessage: problem sending password reset email (%s)", email),
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

	subject := fmt.Sprintf("[ZeroToHero] Hi %s, you've successfully reset your password.", name)

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("RelayPasswordResetConfirmationMessage: problem sending passwrod reset confirmation email (%s)", email),
		)
	}

	return nil
}
