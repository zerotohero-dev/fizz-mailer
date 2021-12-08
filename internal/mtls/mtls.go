package mtls

import "github.com/zerotohero-dev/fizz-mailer/internal/service"



func ListenAndServe(mailerArgs service.Args, spireArgs SpireArgs) {
	runSpireMtlSServer(mailerArgs, spireArgs)
}
