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

const enrolledTpl = `Hi %s,

Welcome 🎉 🐢

I just wanted to let you know that you have successfully subscribed to
FizzBuzz Pro.

You can access to your interview questions, answers, tips, tricks,
study guides and more from https://fizzbuzz.pro/ .

🐢 Can I Ask You a Favor?

I spend a lot of time an energy to make https://fizzbuzz.pro/ unique.

I value quality over quantity, and do my best to squeeze as much 
knowledge and experience as I can into it for you to nail that godforsaken interview.

I merely ask two favors in return:

⚜️ ️Be a FizzBuzz Pro ambassador: Spread the word out, wherever you can…
Be it Facebook, twitter, your friends, your colleagues…
The more people hear about FizzBuzz Pro, the greater of an audience we’ll reach.

⚜️ Send me feedback: If there is something I can add to a solution, or if you
think explaining a particular question differently would be better, tell me about it.

Keep me in the loop, so that I can make the content even better.

I am thrilled to create a quality resource, and I appreciate your help.

You can send your emails to me@volkan.io ✉️.

🐢 Ask Me Anything

You can ask me programming-related, or not.

I have been in the industry for a looong time, and I may have a trick or two
in my sleeves 😉.

May the source be with you.

Volkan.
%s`

type SubscribedMessageParams struct {
	Name string
}

func SubscribedMessageBody(p SubscribedMessageParams) string {
	name := p.Name

	return fmt.Sprintf(enrolledTpl, name, footerTpl)
}
