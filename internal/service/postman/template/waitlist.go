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

const waitlistTpl = `Hello %s,

🐢 Welcome 🎉 🐢

You have successfully joined the waitlist of FizzBuzz Pro.

🐢 🐢 🐢

Stay Tuned!

That’s all about it for now 🙂. 

May the source be with you,

Volkan.
%s`

type WaitlistMessageParams struct {
	Name string
}

func WaitlistMessageBody(p WaitlistMessageParams) string {
	name := p.Name

	return fmt.Sprintf(waitlistTpl, name, footerTpl)
}
