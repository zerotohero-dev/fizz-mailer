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

const welcomeTpl = `Hello %s,

π’ Welcome π π’

You have successfully created your FizzBuzz Pro account.

π’ What to Do Next

π Log in to https://fizzbuzz.pro/ using the email and password youβve signed up with.
π Tap the βSubscribeβ button on your learning dashboard.
π Always be learning.

π’ π’ π’

Thatβs all about it for now π.

May the source be with you,

Volkan.
%s`

type WelcomeMessageParams struct {
	Name string
}

func WelcomeMessageBody(p WelcomeMessageParams) string {
	name := p.Name

	return fmt.Sprintf(welcomeTpl, name, footerTpl)
}
