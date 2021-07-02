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
	"github.com/zerotohero.dev-/fizz-mailer/internal/endpoint"
	"github.com/zerotohero.dev-/fizz-mailer/internal/service"
	"github.com/zerotohero.dev-/fizz-mailer/internal/transport"
)

func InitializeEndpoints(e env.FizzEnv, router *mux.Router) {
	svc := service.New(e, context.Background())

	// Sends email verification email.
	app.Route(
		router, http.NewServer(
			endpoint.MakeSendEmailVerificationEmailEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeSendEmailVerificationEmailRequest),
			app.EncodeResponse,
		),
		"GET", "/v1/send/verification",
	)

	// Sends email verified email.
	app.Route(
		router, http.NewServer(
			endpoint.MakeSendEmailVerifiedEmailEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeSendEmailVerifiedEmailRequest),
			app.EncodeResponse,
		),
		"GET", "/v1/send/verified",
	)

	// Sends welcome email.
	app.Route(
		router, http.NewServer(
			endpoint.MakeSendWelcomeEmailEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeSendWelcomeEmailRequest),
			app.EncodeResponse,
		),
		"GET", "/v1/send/welcome",
	)

	// Sends password reset email.
	app.Route(
		router, http.NewServer(
			endpoint.MakeSendPasswordResetEmailEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeSendPasswordResetEmailRequest),
			app.EncodeResponse,
		),
		"GET", "/v1/send/reset",
	)

	// Sends password reset confirmation email.
	app.Route(
		router, http.NewServer(
			endpoint.MakeSendPasswordResetConfirmationEmailEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeSendPasswordResetConfirmationEmailRequest),
			app.EncodeResponse,
		),
		"GET", "/v1/send/confirm",
	)
}