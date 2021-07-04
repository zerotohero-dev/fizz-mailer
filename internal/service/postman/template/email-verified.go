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

package template

import "fmt"

const emailVerifiedMessageTpl = `
Hello %s,

You have successfully verified your email. You are almost there.

🦄 What to Do Next

To finish you account creation tap the link below and fill in the form
that comes up: 

👉  %s 👈

May the source be with you,

Volkan.

%s`

type EmailVerifiedMessageParams struct {
	Name      string
	SignUpUrl string
}

func EmailVerifiedMessageBody(p EmailVerifiedMessageParams) string {
	name := p.Name
	url := p.SignUpUrl

	return fmt.Sprintf(emailVerifiedMessageTpl, name, url, footerTpl)
}
