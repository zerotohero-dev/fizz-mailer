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
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/fizz-mailer/internal/service"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-mtls/pkg/mtls"
	"net"
)

func runSpireMtlSServer(svcArgs service.Args, spireArgs mtls.SpireArgs) {

//ServerAddress:  e.Crypto.MtlsServerAddress,
//	SocketPath:     e.Spire.SocketPath,
//		AppTrustDomain: e.Spire.AppTrustDomainFizz,
//		AppPrefix:      e.Spire.AppPrefixFizz,
//		AppNameDefault: e.Spire.AppNameFizzDefault,
//		AppName:        e.Crypto.ServiceName,
//		AppNameIdm:     e.Idm.ServiceName,
//		AppNameMailer:  e.Mailer.ServiceName,

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	svc := service.New(svcArgs, ctx)

	handler := func(conn net.Conn) {
		handleConnection(conn, svc)
	}

	poop := func(err error) {
		handleError(err)
	}

	mtls.ListenAndServe(
		svcArgs.MtlsServerAddress,
		svcArgs.MtlsSocketPath,
		svcArgs.MtlsAppName,
		mtls.AllowList(svcArgs.IsDevelopment),
		handler, poop,
	)



	trustDomain := spireArgs.AppTrustDomain
	appPrefix := spireArgs.AppPrefix
	idmAppName := spireArgs.AppNameIdm
	mailerAppName := spireArgs.AppNameMailer
	anyAppName := spireArgs.AppNameDefault

	var ids []spiffeid.ID
	if svcArgs.IsDevelopment {
		anyId, _ := spiffeid.New(trustDomain, appPrefix, anyAppName)
		ids = []spiffeid.ID{anyId}
	} else {
		appId, _ := spiffeid.New(trustDomain, appPrefix, idmAppName)
		mailerId, _ := spiffeid.New(trustDomain, appPrefix, mailerAppName)
		ids = []spiffeid.ID{appId, mailerId}
	}

	log.Info("Crypto mTLS server will try listening… (%s)", spireArgs.ServerAddress)

	listener, err := spiffetls.ListenWithMode(ctx, "tcp", spireArgs.ServerAddress,
		spiffetls.MTLSServerWithSourceOptions(
			tlsconfig.AuthorizeOneOf(ids...),
			workloadapi.WithClientOptions(workloadapi.WithAddr(spireArgs.SocketPath)),
		))

	if err != nil {
		log.Err("runSpireMtlSServer: Unable to create TLS listener: %v", err.Error())
		panic(err.Error())
	}

	log.Info(
		"🐢 Service '%s' is waiting for mTLS connections at '%s",
		spireArgs.AppName, spireArgs.ServerAddress,
	)

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			panic(err.Error())
		}
	}(listener)



	for {
		conn, err := listener.Accept()
		if err != nil {
			go handleError(err)
		}
		go
	}
}
