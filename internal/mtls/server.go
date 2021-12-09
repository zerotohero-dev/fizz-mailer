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
	"context"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-mailer/internal/service"
	"github.com/zerotohero-dev/fizz-mtls/pkg/mtls"
	"github.com/zerotohero-dev/fizz-mtls/pkg/mtls/ext"
	"net"
)

func ListenAndServe(svcArgs service.Args, spireArgs ext.SpireArgs) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	svc := service.New(svcArgs, ctx)

	mux := func(conn net.Conn) {
		handleConnection(conn, svc)
	}

	allowedIds, err := ext.AllowList(spireArgs, svcArgs.IsDevelopment)
	if err != nil {
		log.Err("ListenAndServe: Unable to acquire SVIDs: %v", err.Error())
	}

	mtls.ListenAndServe(
		svcArgs.MtlsServerAddress, svcArgs.MtlsSocketPath, svcArgs.MtlsAppName,
		allowedIds, mux, handleError,
	)
}
