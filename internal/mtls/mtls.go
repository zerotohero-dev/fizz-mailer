package mtls

import "github.com/zerotohero-dev/fizz-mailer/internal/service"

type SpireArgs struct {
	ServerAddress  string
	SocketPath     string
	AppPrefix      string
	AppNameDefault string
	AppName        string
	AppNameIdm     string
	AppNameMailer  string
	AppTrustDomain string
}

func ListenAndServe(mailerArgs service.Args, spireArgs SpireArgs) {
	runSpireMtlSServer(mailerArgs, spireArgs)
}
