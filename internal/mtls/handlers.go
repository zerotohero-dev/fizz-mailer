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
	"encoding/json"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-mailer/internal/service"
	"github.com/zerotohero-dev/fizz-mtls/pkg/mtls/ext"
	"net"
)

func handleRelayWelcome(conn net.Conn, svc service.Service, body string) {
	body, err := ext.Payload(body)
	if err != nil {
		log.Err("handleRelayWelcome: problem with unmarshal: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelayWelcomeMessageResponse{
			Success: false,
			Err:     "handleRelayWelcome: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleRelayWelcome: could not send: %s", sendErr.Error())
		}
		return
	}

	var req reqres.RelayWelcomeMessageRequest
	err = json.Unmarshal([]byte(body), &req)
	if err != nil {
		log.Err("handleRelayWelcome: problem with unmarshal: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelayWelcomeMessageResponse{
			Success: false,
			Err:     "handleRelayWelcome: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleRelayWelcome: could not send: %s", sendErr.Error())
		}
		return
	}

	err = svc.RelayWelcomeMessage(req.Email, req.Name)
	if err != nil {
		log.Err("handleRelayWelcome: problem with relaying message: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelayWelcomeMessageResponse{
			Success: false,
		})
		if sendErr != nil {
			log.Err("handleRelayWelcome: could not send: %s", sendErr.Error())
		}
		return
	}

	sendErr := ext.Send(conn, reqres.RelayWelcomeMessageResponse{
		Success: true,
	})
	if sendErr != nil {
		log.Err("handleRelayWelcome: could not send: %s", sendErr.Error())
	}
}

func handleRelayEmailVerification(conn net.Conn, svc service.Service, body string) {
	body, err := ext.Payload(body)
	if err != nil {
		log.Err("handleRelayEmailVerification: problem with unmarshal: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelayEmailVerificationMessageResponse{
			Success: false,
			Err:     "handleRelayEmailVerification: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleRelayEmailVerification: could not send: %s", sendErr.Error())
		}
		return
	}

	var req reqres.RelayEmailVerificationMessageRequest
	err = json.Unmarshal([]byte(body), &req)
	if err != nil {
		log.Err("handleRelayEmailVerification: problem with unmarshal: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelayEmailVerificationMessageResponse{
			Success: false,
			Err:     "handleRelayEmailVerification: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleRelayEmailVerification: could not send: %s", sendErr.Error())
		}
		return
	}

	err = svc.RelayEmailVerificationMessage(req.Email, req.Name, req.Token)
	if err != nil {
		log.Err("handleRelayEmailVerification: problem with relaying message: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelayEmailVerificationMessageResponse{
			Success: false,
		})
		if sendErr != nil {
			log.Err("handleRelayEmailVerification: could not send: %s", sendErr.Error())
		}
		return
	}

	sendErr := ext.Send(conn, reqres.RelayEmailVerificationMessageResponse{
		Success: true,
	})
	if sendErr != nil {
		log.Err("handleRelayEmailVerification: could not send: %s", sendErr.Error())
	}
}

func handleRelayPasswordReset(conn net.Conn, svc service.Service, body string) {
	body, err := ext.Payload(body)
	if err != nil {
		log.Err("handleRelayPasswordReset: problem with unmarshal: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelayPasswordResetMessageResponse{
			Success: false,
			Err:     "handleRelayPasswordReset: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleRelayPasswordReset: could not send: %s", sendErr.Error())
		}
		return
	}

	var req reqres.RelayPasswordResetMessageRequest
	err = json.Unmarshal([]byte(body), &req)
	if err != nil {
		log.Err("handleRelayPasswordReset: problem with unmarshal: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelayPasswordResetMessageResponse{
			Success: false,
			Err:     "handleRelayPasswordReset: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleRelayPasswordReset: could not send: %s", sendErr.Error())
		}
		return
	}

	err = svc.RelayPasswordResetMessage(req.Email, req.Name, req.Token)
	if err != nil {
		log.Err("handleRelayPasswordReset: problem with relaying message: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelayPasswordResetMessageResponse{
			Success: false,
		})
		if sendErr != nil {
			log.Err("handleRelayPasswordReset: could not send: %s", sendErr.Error())
		}
		return
	}

	sendErr := ext.Send(conn, reqres.RelayPasswordResetMessageResponse{
		Success: true,
	})
	if sendErr != nil {
		log.Err("handleRelayPasswordReset: could not send: %s", sendErr.Error())
	}
}

func handleRelayPasswordResetConfirm(conn net.Conn, svc service.Service, body string) {
	body, err := ext.Payload(body)
	if err != nil {
		log.Err("handleRelayPasswordResetConfirm: problem with unmarshal: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelayPasswordResetConfirmationMessageResponse{
			Success: false,
			Err:     "handleRelayPasswordResetConfirm: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleRelayPasswordResetConfirm: could not send: %s", sendErr.Error())
		}
		return
	}

	var req reqres.RelayPasswordResetConfirmationMessageRequest
	err = json.Unmarshal([]byte(body), &req)
	if err != nil {
		log.Err("handleRelayPasswordResetConfirm: problem with unmarshal: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelayPasswordResetConfirmationMessageResponse{
			Success: false,
			Err:     "handleRelayPasswordResetConfirm: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleRelayPasswordResetConfirm: could not send: %s", sendErr.Error())
		}
		return
	}

	err = svc.RelayPasswordResetConfirmationMessage(req.Email, req.Name)
	if err != nil {
		log.Err("handleRelayPasswordResetConfirm: problem with relaying message: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelayPasswordResetConfirmationMessageResponse{
			Success: false,
		})
		if sendErr != nil {
			log.Err("handleRelayPasswordResetConfirm: could not send: %s", sendErr.Error())
		}
		return
	}

	sendErr := ext.Send(conn, reqres.RelayPasswordResetConfirmationMessageResponse{
		Success: true,
	})
	if sendErr != nil {
		log.Err("handleRelayPasswordResetConfirm: could not send: %s", sendErr.Error())
	}
}

func handleRelaySubscriptionConfirm(conn net.Conn, svc service.Service, body string) {
	body, err := ext.Payload(body)
	if err != nil {
		log.Err("handleRelaySubscriptionConfirm: problem with unmarshal: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelaySubscribedMessageResponse{
			Success: false,
			Err:     "handleRelaySubscriptionConfirm: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleRelaySubscriptionConfirm: could not send: %s", sendErr.Error())
		}
		return
	}

	var req reqres.RelaySubscribedMessageRequest
	err = json.Unmarshal([]byte(body), &req)
	if err != nil {
		log.Err("handleRelaySubscriptionConfirm: problem with unmarshal: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelaySubscribedMessageResponse{
			Success: false,
			Err:     "handleRelaySubscriptionConfirm: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleRelaySubscriptionConfirm: could not send: %s", sendErr.Error())
		}
		return
	}

	err = svc.RelaySubscribedMessage(req.Email, req.Name)
	if err != nil {
		log.Err("handleRelaySubscriptionConfirm: problem with relaying message: %s", err.Error())
		sendErr := ext.Send(conn, reqres.RelaySubscribedMessageResponse{
			Success: false,
		})
		if sendErr != nil {
			log.Err("handleRelaySubscriptionConfirm: could not send: %s", sendErr.Error())
		}
		return
	}

	sendErr := ext.Send(conn, reqres.RelaySubscribedMessageResponse{
		Success: true,
	})
	if sendErr != nil {
		log.Err("handleRelaySubscriptionConfirm: could not send: %s", sendErr.Error())
	}
}

func handleUnknown(conn net.Conn, svc service.Service, body string) {
	// Unknown request. Better to close the connection as fast as possible.
	if sendErr := ext.Send(conn, ""); sendErr != nil {
		log.Err("handleUnknown: could not send: %s", sendErr.Error())
	}
}
