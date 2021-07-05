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

import "github.com/zerotohero-dev/fizz-mailer/internal/service/postman"

// For development environments, this service logs the email body to
// the standard output instead of relaying the email. This is by design,
// as we want to have a limited “allow list” for the backend servers only.
// Plus it speeds up development time, since sometimes it takes a while for
// those emails to arrive. You can configure this service to run in production
// mode and also add your developer machine’s IP to the allow list, if you want
// to have a more production-like behavior.

func (s service) RelayEmailVerificationMessage(email, name, emailVerificationToken string) error {
	return postman.RelayEmailVerificationMessage(s.env, email, name, emailVerificationToken)
}

func (s service) RelayEmailVerifiedMessage(email, name string) error {
	return postman.RelayEmailVerifiedMessage(s.env, email, name)
}

func (s service) RelayWelcomeMessage(email, name string) error {
	return postman.RelayWelcomeMessage(s.env, email, name)
}

func (s service) RelayPasswordResetMessage(email, name, passwordResetToken string) error {
	return postman.RelayPasswordResetMessage(s.env, email, name, passwordResetToken)
}

func (s service) RelayPasswordResetConfirmationMessage(email, name string) error {
	return postman.RelayPasswordResetConfirmationMessage(s.env, email, name)
}

func (s service) RelaySubscribedMessage(email, name string) error {
	return postman.RelaySubscribedMessage(s.env, email, name)
}
