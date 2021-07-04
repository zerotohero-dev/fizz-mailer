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

func MakeRelayEmailVerifiedMessageEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		if gr, hasProblem := request.(reqres.ContentTypeProblemRequest); hasProblem {
			return reqres.RelayEmailVerifiedMessageResponse{
				Success: false,
				Err:     gr.Err,
			}, nil
		}

		req := request.(reqres.RelayEmailVerifiedMessageRequest)
		if req.Err != "" {
			return reqres.RelayEmailVerifiedMessageResponse{
				Success: false,
				Err:     req.Err,
			}, nil
		}

		email := sanitization.SanitizeEmail(req.Email)
		if email == "" {
			return reqres.RelayEmailVerifiedMessageResponse{
				Success: false,
				Err:     "Error: missing email",
			}, nil
		}

		fullName := sanitization.SanitizeName(req.Name)
		if fullName == "" {
			return reqres.RelayEmailVerifiedMessageResponse{
				Success: false,
				Err:     "Error: missing full name",
			}, nil
		}

		err := svc.RelayEmailVerifiedMessage(email, fullName)
		if err != nil {
			return reqres.RelayEmailVerifiedMessageResponse{
				Success: false,
				Err:     "Error sending user verification email",
			}, nil
		}

		return reqres.RelayEmailVerifiedMessageResponse{
			Success: true,
		}, nil
	}
}