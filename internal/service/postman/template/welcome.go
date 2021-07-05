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

🐢 Welcome 🎉 🐢

You have successfully created your FizzBuzz Pro account.

🐢 What to Do Next

👉 Log in to https://fizzbuzz.pro/ using the email and password you’ve signed up with.
👉 Tap the “Subscribe” button on your learning dashboard.
👉 Always be learning.

🐢 🐢 🐢

That’s all about it for now 🙂.

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
