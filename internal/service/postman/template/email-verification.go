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

import (
	"fmt"
	"net/url"
)

const activateTpl = `Hello %s,

To sign up to your FizzBuzz pro account, you need to verify your email first.

Follow this link and you’ll be all set in no time:

%s

(if you cannot click the link, copy and paste it to your browser’s address bar)

May the source be with you,

Volkan.
%s`

type EmailVerificationMessageParams struct {
	Name       string
	Email      string
	Token      string
	IdmBaseUrl string
}

func EmailVerificationMessageBody(p EmailVerificationMessageParams) string {
	name := p.Name
	email := url.QueryEscape(p.Email)
	token := url.QueryEscape(p.Token)
	verificationUrl := fmt.Sprintf(
		"%s/verify?email=%s&token=%s",
		p.IdmBaseUrl,
		email,
		token,
	)

	return fmt.Sprintf(activateTpl, name, verificationUrl, footerTpl)
}
