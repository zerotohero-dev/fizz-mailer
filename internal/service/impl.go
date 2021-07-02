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
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
)

// For development environments, this service logs the email body to
// the standard output instead of relaying the email. This is by design,
// as we want to have a limited “allow list” for the backend servers only.
// Plus it speeds up development time, since sometimes it takes a while for
// those emails to arrive. You can configure this service to run in production
// mode and also add your developer machine’s IP to the allow list, if you want
// to have a more production-like behavior.

func (s service) RelayEmailVerificationMessage(email, name, emailVerificationToken string) error {
	if s.env.Deployment.Type == env.Development {
		log.Info("mailer:dev:RelayEmailVerificationMessage")

		return nil
	}

	return nil
}

func (s service) RelayEmailVerifiedMessage(email, name string) error {
	if s.env.Deployment.Type == env.Development {
		log.Info("mailer:dev:RelayEmailVerifiedMessage")

		return nil
	}

	return nil
}

func (s service) RelayWelcomeMessage(email, name string) error {
	if s.env.Deployment.Type == env.Development {
		log.Info("mailer:dev:RelayWelcomeMessage")

		return nil
	}

	return nil
}

func (s service) RelayPasswordResetMessage(email, name, passwordResetToken string) error {
	if s.env.Deployment.Type == env.Development {
		log.Info("mailer:dev:RelayPasswordResetMessage")

		return nil
	}

	return nil
}

func (s service) RelayPasswordResetConfirmationMessage(email, name string) error {
	if s.env.Deployment.Type == env.Development {
		log.Info("mailer:dev:RelayPasswordResetConfirmationMessage")

		return nil
	}

	return nil
}