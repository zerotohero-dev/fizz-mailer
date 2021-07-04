package postman

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"time"
)

func RelayEmailVerificationMessage(e env.FizzEnv, email, name, emailVerificationToken string) error {
	body := template.ActivateEmailBody(template.ActivateEmailParams{
		Email: email,
		Alias: name,
		Token: emailVerificationToken,
	})

	domain := env.Env.MailgunDomain
	apiKey := env.Env.MailgunApiKey

	subject := fmt.Sprintf("[ZeroToHero] %s, please verify your account 🦄", name)

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("SendUserActivationEmail: problem sending activation email (%s)", email),
		)
	}

	return nil
}

func RelayEmailVerifiedMessage(e env.FizzEnv, email, name string) error {
	body := template.AccountVerifiedEmailBody(template.AccountVerifiedParams{
		Alias: name,
	})

	domain := env.Env.MailgunDomain
	apiKey := env.Env.MailgunApiKey

	subject := fmt.Sprintf("[ZeroToHero] %s, you have activated your account 🦄", name)

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("SendUserVerifiedEmail: problem account verification confirmation email (%s)", email),
		)
	}

	return nil
}

func RelayWelcomeMessage(e env.FizzEnv, email, name string) error {
	body := template.EnrolledEmailBody(template.EnrolledEmailParams{
		Alias:  name,
		Course: course,
	})

	domain := env.Env.MailgunDomain
	apiKey := env.Env.MailgunApiKey

	subject := fmt.Sprintf("[ZeroToHero] %s, Welcome to %s 🦄", name, course)

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("SendUserWelcomeEmail: problem sending welcome email (%s)", email),
		)
	}

	return nil
}

func RelayPasswordResetMessage(e env.FizzEnv, email, name, passwordResetToken string) error {
	body := template.PasswordResetEmailBody(template.PasswordResetEmailParams{
		Alias: name,
		Email: email,
		Token: passwordResetToken,
	})

	domain := env.Env.MailgunDomain
	apiKey := env.Env.MailgunApiKey

	subject := fmt.Sprintf("[ZeroToHero] Hi %s, here are your password reset instructions.", name)

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("SendUserPasswordResetEmail: problem sending password reset email (%s)", email),
		)
	}

	return nil
}

func RelayPasswordResetConfirmationMessage(e env.FizzEnv, email, name string) error {
	body := template.PasswordResetConfirmationEmailBody(template.PasswordResetConfirmationEmailParams{
		Alias: name,
	})

	domain := env.Env.MailgunDomain
	apiKey := env.Env.MailgunApiKey

	subject := fmt.Sprintf("[ZeroToHero] Hi %s, you've successfully reset your password.", name)

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(from, subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, _, err := mg.Send(ctx, m)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("SendUserWelcomeEmail: problem sending passwrod reset confirmation email (%s)", email),
		)
	}

	return nil
}
