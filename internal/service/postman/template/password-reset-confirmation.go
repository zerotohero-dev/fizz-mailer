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
)

const passwordResetConfirmationTpl = `Hello %s,

You’ve just reset your password.

If that was not you, please let me know by replying to this email.

Please let me know if you have any other issues.

May the source be with you,

Volkan.
%s`

type PasswordResetConfirmationMessageParams struct {
	Name string
}

func PasswordResetConfirmationMessageBody(p PasswordResetConfirmationMessageParams) string {
	name := p.Name

	return fmt.Sprintf(passwordResetConfirmationTpl, name, footerTpl)
}
