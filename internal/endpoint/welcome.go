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

package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-mailer/internal/service"
	"github.com/zerotohero-dev/fizz-validation/pkg/sanitization"
)

func MakeRelayWelcomeMessageEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		if gr, hasProblem := request.(reqres.ContentTypeProblemRequest); hasProblem {
			return reqres.RelayEmailVerificationMessageResponse{
				Success: false,
				Err:     gr.Err,
			}, nil
		}

		req := request.(reqres.RelayEmailVerificationMessageRequest)
		if req.Err != "" {
			return reqres.RelayEmailVerificationMessageResponse{
				Success: false,
				Err:     req.Err,
			}, nil
		}

		email := sanitization.SanitizeEmail(req.Email)
		if email == "" {
			return reqres.RelayEmailVerificationMessageResponse{
				Success: false,
				Err:     "Error: missing email",
			}, nil
		}

		fullName := sanitization.SanitizeName(req.Name)
		if fullName == "" {
			return reqres.RelayEmailVerificationMessageResponse{
				Success: false,
				Err:     "Error: missing full name",
			}, nil
		}

		err := svc.RelayWelcomeMessage(email, fullName)
		if err != nil {
			return reqres.RelayEmailVerificationMessageResponse{
				Success: false,
				Err:     "Error sending user welcome email",
			}, nil
		}

		return reqres.RelayEmailVerificationMessageResponse{
			Success: true,
		}, nil
	}
}