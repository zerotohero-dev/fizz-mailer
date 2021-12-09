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

package main

import (
	"github.com/gorilla/mux"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-mailer/internal/mtls"
	"github.com/zerotohero-dev/fizz-mailer/internal/service"
	"github.com/zerotohero-dev/fizz-mtls/pkg/mtls/ext"
)

func listenAndServeApp(e env.FizzEnv) {
	go func() {
		svcName := e.Mailer.ServiceName

		r := mux.NewRouter()

		app.RouteHealthEndpoints(e.Mailer.PathPrefix, r)
		app.ListenAndServe(e, svcName, e.Mailer.Port, r)
	}()
}

func listenAndServeMtls(e env.FizzEnv) {
	svcName := e.Mailer.ServiceName

	mtls.ListenAndServe(service.Args{
		IsDevelopment:     e.Deployment.Type == env.Development,
		MtlsServerAddress: e.Mailer.MtlsServerAddress,
		MtlsSocketPath:    e.Spire.SocketPath,
		MtlsAppName:       svcName,
	},
		ext.SpireArgs{
			AppTrustDomain: e.Spire.AppTrustDomainFizz,
			AppPrefix:      e.Spire.AppPrefixFizz,
			AppNameDefault: e.Spire.AppNameFizzDefault,
			AppName:        svcName,
			AppNameIdm:     e.Idm.ServiceName,
			AppNameMailer:  svcName,
		},
	)
}
