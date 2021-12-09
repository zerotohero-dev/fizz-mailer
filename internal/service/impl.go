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

func args(s service, email, name string) postman.Args {
	return postman.Args{
		MailgunDomain: s.args.MailgunDomain,
		MailgunApiKey: s.args.MailgunApiKey,
		IsDevelopment: s.args.IsDevelopment,
		Email:         email,
		Name:          name,
	}
}

func (s service) RelayEmailVerificationMessage(email, name, emailVerificationToken string) error {
	return postman.RelayEmailVerificationMessage(
		args(s, email, name),
		s.args.EmailVerificationBaseUrl,
		emailVerificationToken,
	)
}

func (s service) RelayWelcomeMessage(email, name string) error {
	return postman.RelayWelcomeMessage(args(s, email, name))
}

func (s service) RelayPasswordResetMessage(email, name, passwordResetToken string) error {
	return postman.RelayPasswordResetMessage(
		args(s, email, name),
		s.args.PasswordResetBaseUrl,
		passwordResetToken,
	)
}

func (s service) RelayPasswordResetConfirmationMessage(email, name string) error {
	return postman.RelayPasswordResetConfirmationMessage(
		args(s, email, name),
	)
}

func (s service) RelaySubscribedMessage(email, name string) error {
	return postman.RelaySubscribedMessage(
		args(s, email, name),
	)
}
