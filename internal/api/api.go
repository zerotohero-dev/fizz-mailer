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

package api

import (
	"context"
	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-mailer/internal/endpoint"
	"github.com/zerotohero-dev/fizz-mailer/internal/service"
	"github.com/zerotohero-dev/fizz-mailer/internal/transport"
)

func InitializeEndpoints(e env.FizzEnv, router *mux.Router) {
	svc := service.New(e, context.Background())

	prefix := e.Mailer.PathPrefix

	// Sends email verification email.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeRelayEmailVerificationMessageEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeRelayEmailVerificationMessageRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, "/v1/relay/verification",
	)

	// Sends email verified email.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeRelayEmailVerifiedMessageEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeRelayEmailVerifiedMessageRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, "/v1/relay/verified",
	)

	// Sends welcome email.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeRelayWelcomeMessageEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeRelayWelcomeMessageRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, "/v1/relay/welcome",
	)

	// Sends password reset email.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeRelayPasswordResetMessageEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeRelayPasswordResetMessageRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, "/v1/relay/reset",
	)

	// Sends password reset confirmation email.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeRelayPasswordResetConfirmationMessageEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeRelayPasswordResetConfirmationMessageRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, "/v1/relay/confirm",
	)

	// Sends subscription confirmation.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeRelaySubscribedMessageEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeRelaySubscribedMessageRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, "/v1/relay/subscribed",
	)
}
