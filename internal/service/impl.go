/*
 *  \
 *  \\,
 *   \\\,^,.,,.                    “Zero to Hero”
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package service

func (s service) SendEmailVerificationEmail(email, name, emailVerificationToken string) error {
	panic("implement me")
}

func (s service) SendEmailVerifiedEmail(email, name string) error {
	panic("implement me")
}

func (s service) SendWelcomeEmail(email, name string) error {
	panic("implement me")
}

func (s service) SendPasswordResetEmail(email, name, passwordResetToken string) error {
	panic("implement me")
}

func (s service) SendPasswordResetConfirmationEmail(email, name string) error {
	panic("implement me")
}