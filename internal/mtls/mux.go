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
	apiEndpoint string, apiMethod method.Method, conn net.Conn, svc service.Service,
) {
	switch {
	case apiEndpoint == endpoint.Mailer.Verification && apiMethod == method.Post:
		_ = handleRelayEmailVerification(conn, svc)
	case apiEndpoint == endpoint.Crypto.Jwt && apiMethod == method.Post:
		_ = handleCryptoJwt(conn, svc)
	case apiEndpoint == endpoint.Crypto.SecureHash && apiMethod == method.Post:
		_ = handleSecureHash(conn, svc)
	case apiEndpoint == endpoint.Crypto.SecureToken && apiMethod == method.Get:
		_ = handleSecureToken(conn, svc)
	default:
		_ = handleUnknown(conn, svc)
	}
}
