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

package mtls

import (
	"github.com/zerotohero-dev/fizz-entity/pkg/endpoint"
	"github.com/zerotohero-dev/fizz-entity/pkg/method"
	"github.com/zerotohero-dev/fizz-mailer/internal/service"
	"net"
)

func multiplex(
	apiEndpoint string, apiMethod method.Method, body string,
	conn net.Conn, svc service.Service,
) {
	switch {
	case apiEndpoint == endpoint.Mailer.Verification &&
		apiMethod == method.Post:
		handleRelayEmailVerification(conn, svc, body)
	case apiEndpoint == endpoint.Mailer.Welcome &&
		apiMethod == method.Post:
		handleRelayWelcome(conn, svc, body)
	case apiEndpoint == endpoint.Mailer.PasswordReset &&
		apiMethod == method.Post:
		handleRelayPasswordReset(conn, svc, body)
	case apiEndpoint == endpoint.Mailer.PasswordResetConfirm &&
		apiMethod == method.Get:
		handleRelayPasswordResetConfirm(conn, svc, body)
	case apiEndpoint == endpoint.Mailer.SubscriptionConfirm &&
		apiMethod == method.Get:
		handleRelaySubscriptionConfirm(conn, svc, body)
	default:
		handleUnknown(conn, svc, body)
	}
}
