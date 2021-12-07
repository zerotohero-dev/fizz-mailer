/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package service

import (
	"context"
)

type Service interface {
	RelayEmailVerificationMessage(email, name, emailVerificationToken string) error
	RelayWelcomeMessage(email, name string) error
	RelayPasswordResetMessage(email, name, passwordResetToken string) error
	RelayPasswordResetConfirmationMessage(email, name string) error
	RelaySubscribedMessage(email, name string) error
}

type Args struct {
	IsDevelopment bool
	MailgunDomain string
	MailgunApiKey string
	EmailVerificationBaseUrl string
	PasswordResetBaseUrl string
}

type service struct {
	args Args
	ctx context.Context
}

func New(args Args, ctx context.Context) Service {
	return &service{
		args: args,
		ctx: ctx,
	}
}
