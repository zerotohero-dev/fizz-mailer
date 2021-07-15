#!/usr/bin/env zsh

#!/usr/bin/env zsh

#  \
#  \\,
#   \\\,^,.,,.                     Zero to Hero
#   ,;7~((\))`;;,,               <zerotohero.dev>
#   ,(@') ;)`))\;;',    stay up to date, be curious: learn
#    )  . ),((  ))\;,
#   /;`,,/7),)) )) )\,,
#  (& )`   (,((,((;( ))\,

TAG="0.0.10"

# shellcheck disable=SC1090
source ~/.zprofile

docker run -e FIZZ_DEPLOYMENT_TYPE="$FIZZ_DEPLOYMENT_TYPE" \
-e FIZZ_MAILER_SVC_PORT="$FIZZ_MAILER_SVC_PORT" \
-e FIZZ_MAILER_EMAIL_VERIFICATION_BASE_URL="$FIZZ_MAILER_EMAIL_VERIFICATION_BASE_URL" \
-e FIZZ_MAILER_WEB_APP_HOST_BASE_URL="$FIZZ_MAILER_WEB_APP_HOST_BASE_URL" \
-e FIZZ_MAILER_EMAIL_VERIFICATION_BASE_URL="$FIZZ_MAILER_EMAIL_VERIFICATION_BASE_URL" \
-e FIZZ_MAILER_PASSWORD_RESET_BASE_URL="$FIZZ_MAILER_PASSWORD_RESET_BASE_URL" \
-e FIZZ_MAILER_HONEYBADGER_API_KEY="$FIZZ_MAILER_HONEYBADGER_API_KEY" \
-e FIZZ_MAILER_MAILGUN_DOMAIN="$FIZZ_MAILER_MAILGUN_DOMAIN" \
-e FIZZ_MAILER_MAILGUN_API_KEY="$FIZZ_MAILER_MAILGUN_API_KEY"
zerotohero-dev/fizz-mailer:$TAG
