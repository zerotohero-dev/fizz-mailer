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



func (s service) RelayEmailVerificationMessage(email, name, emailVerificationToken string) error {
	return postman.RelayEmailVerificationMessage(s.args, email, name, emailVerificationToken)
}

func (s service) RelayWelcomeMessage(email, name string) error {
	return postman.RelayWelcomeMessage(s.args, email, name)
}

func (s service) RelayPasswordResetMessage(email, name, passwordResetToken string) error {
	return postman.RelayPasswordResetMessage(s.args, email, name, passwordResetToken)
}

func (s service) RelayPasswordResetConfirmationMessage(email, name string) error {
	return postman.RelayPasswordResetConfirmationMessage(s.args, email, name)
}

func (s service) RelaySubscribedMessage(email, name string) error {
	return postman.RelaySubscribedMessage(s.args, email, name)
}
