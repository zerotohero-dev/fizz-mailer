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

var urls = struct {
	EmailVerification    string
	Welcome              string
	PasswordReset        string
	PasswordResetConfirm string
	SubscriptionConfirm  string
}{
	EmailVerification:    "/v1/relay/verification",
	Welcome:              "/v1/relay/welcome",
	PasswordReset:        "/v1/relay/reset",
	PasswordResetConfirm: "/v1/relay/confirm",
	SubscriptionConfirm:  "/v1/relay/subscribed",
}

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
		router, "POST", prefix, urls.EmailVerification,
	)

	// Sends welcome email.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeRelayWelcomeMessageEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeRelayWelcomeMessageRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, urls.Welcome,
	)

	// Sends password reset email.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeRelayPasswordResetMessageEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeRelayPasswordResetMessageRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, urls.PasswordReset,
	)

	// Sends password reset confirmation email.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeRelayPasswordResetConfirmationMessageEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeRelayPasswordResetConfirmationMessageRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, urls.PasswordResetConfirm,
	)

	// Sends subscription confirmation.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeRelaySubscribedMessageEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeRelaySubscribedMessageRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, urls.SubscriptionConfirm,
	)
}
