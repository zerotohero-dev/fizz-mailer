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
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

func main() {
	e := *env.New()

	configureApp(e)
	listenAndServeApp(e)
	listenAndServeMtls(e)
}

//func main() {
//	e := *env.New()
//
//	appEnv := e.Mailer
//	svcName := appEnv.ServiceName
//
//	app.Configure(e, svcName, appEnv.HoneybadgerApiKey, appEnv.Sanitize)
//
//	r := mux.NewRouter()
//	api.InitializeEndpoints(e, r)
//	app.RouteHealthEndpoints(e.Mailer.PathPrefix, r)
//
//	app.ListenAndServe(e, svcName, appEnv.Port, r)
//}
