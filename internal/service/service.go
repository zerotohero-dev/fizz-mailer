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
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

type Service interface {
	RelayEmailVerificationMessage(email, name, emailVerificationToken string) error
	RelayEmailVerifiedMessage(email, name string) error
	RelayWelcomeMessage(email, name string) error
	RelayPasswordResetMessage(email, name, passwordResetToken string) error
	RelayPasswordResetConfirmationMessage(email, name string) error
	RelaySubscribedMessage(email, name string) error
}

type service struct {
	env env.FizzEnv
	ctx context.Context
}

func (s service) RelaySubscribedMessage(email, name string) error {
	panic("implement me")
}

func New(e env.FizzEnv, ctx context.Context) Service {
	return &service{
		env: e,
		ctx: ctx,
	}
}
