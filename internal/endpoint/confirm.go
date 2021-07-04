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

func MakeRelayPasswordResetConfirmationMessageEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		if gr, hasProblem := request.(reqres.ContentTypeProblemRequest); hasProblem {
			return reqres.RelayPasswordResetConfirmationMessageResponse{
				Success: false,
				Err:     gr.Err,
			}, nil
		}

		req := request.(reqres.RelayPasswordResetConfirmationMessageRequest)
		if req.Err != "" {
			return reqres.RelayPasswordResetConfirmationMessageResponse{
				Success: false,
				Err:     req.Err,
			}, nil
		}

		email := sanitization.SanitizeEmail(req.Email)
		if email == "" {
			return reqres.RelayPasswordResetConfirmationMessageResponse{
				Success: false,
				Err:     "Error: missing email",
			}, nil
		}

		alias := sanitization.SanitizeName(req.Name)
		if alias == "" {
			return reqres.RelayPasswordResetConfirmationMessageResponse{
				Success: false,
				Err:     "Error: missing alias",
			}, nil
		}

		err := svc.RelayPasswordResetConfirmationMessage(email, alias)
		if err != nil {
			return reqres.RelayPasswordResetConfirmationMessageResponse{
				Success: false,
				Err:     "Error sending user welcome email",
			}, nil
		}

		return reqres.RelayPasswordResetConfirmationMessageResponse{
			Success: true,
		}, nil
	}
}