package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-mailer/internal/service"
	"github.com/zerotohero-dev/fizz-validation/pkg/sanitization"
)

func MakeRelaySubscribedMessageEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		if gr, hasProblem := request.(reqres.ContentTypeProblemRequest); hasProblem {
			return reqres.RelaySubscribedMessageResponse{
				Success: false,
				Err:     gr.Err,
			}, nil
		}

		req := request.(reqres.RelaySubscribedMessageRequest)
		if req.Err != "" {
			return reqres.RelaySubscribedMessageResponse{
				Success: false,
				Err:     req.Err,
			}, nil
		}

		email := sanitization.SanitizeEmail(req.Email)
		if email == "" {
			return reqres.RelaySubscribedMessageResponse{
				Success: false,
				Err:     "Error: missing email",
			}, nil
		}

		alias := sanitization.SanitizeName(req.Name)
		if alias == "" {
			return reqres.RelaySubscribedMessageResponse{
				Success: false,
				Err:     "Error: missing alias",
			}, nil
		}

		err := svc.RelaySubscribedMessage(email, alias)
		if err != nil {
			return reqres.RelaySubscribedMessageResponse{
				Success: false,
				Err:     "Error sending user subscription confirmation email",
			}, nil
		}

		return reqres.RelaySubscribedMessageResponse{
			Success: true,
		}, nil
	}
}
