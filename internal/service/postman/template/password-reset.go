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

const forgotPasswordTpl = `Hello %s,

You just requested to reset your password.

If that was not you, please let me know by replying to this email.

To reset your password, click this link and follow the instructions there:

%s

Please let me know if you have any other issues.

May the source be with you,

Volkan.
%s`

type PasswordResetMessageParams struct {
	Name                 string
	Email                string
	Token                string
	PasswordResetBaseUrl string
}

func PasswordResetMessageBody(p PasswordResetMessageParams) string {
	alias := p.Name
	email := url.QueryEscape(p.Email)
	token := url.QueryEscape(p.Token)

	baseUrl := p.PasswordResetBaseUrl
	passwordResetUrl := fmt.Sprintf(
		"%s?email=%s&token=%s",
		baseUrl,
		email,
		token,
	)

	return fmt.Sprintf(forgotPasswordTpl, alias, passwordResetUrl, footerTpl)
}
