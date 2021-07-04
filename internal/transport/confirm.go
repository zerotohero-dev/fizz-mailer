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

package transport

import (
	"context"
	"encoding/json"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"net/http"
)

func DecodeRelayPasswordResetConfirmationMessageRequest(
	_ context.Context, r *http.Request) (interface{}, error) {
	var request reqres.RelayPasswordResetConfirmationMessageResponse

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Err("DecodeRelayPasswordResetConfirmationMessageRequest: error decoding: %s", err.Error())

		request.Err = "DecodeRelayPasswordResetConfirmationMessageRequest: Problem decoding JSON."
	}

	return request, nil
}