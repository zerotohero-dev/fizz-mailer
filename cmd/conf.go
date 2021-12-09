package main

import (
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

func configureApp(e env.FizzEnv) {
	app.Configure(app.ConfigureOptions{
		AppName:           e.Mailer.ServiceName,
		DeploymentType:    e.Deployment.Type,
		HoneybadgerApiKey: e.Mailer.HoneybadgerApiKey,
		LogDestination:    e.Log.Destination,
		SanitizeFn: func() {
			e.Mailer.Sanitize()
			e.Log.Sanitize()
		},
	})
}
