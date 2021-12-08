```text
 \
 \\,
  \\\,^,.,,.                    “Zero to Hero”
  ,;7~((\))`;;,,               <zerotohero.dev>
  ,(@') ;)`))\;;',    stay up to date, be curious: learn
   )  . ),((  ))\;,
  /;`,,/7),)) )) )\,,
 (& )`   (,((,((;( ))\,
```

## fizz-mailer

## About

`fizz-mailer` relays templated emails using `mailgun`.

## Development Environment Changes

For development environments, this service logs the email body to
the standard output instead of relaying the email. This is by design,
as we want to have a limited “*allow list*” for the backend servers only.
In addition, it speeds up development time, since sometimes it takes a while for
those emails to arrive. 

You can configure this service to run in production mode and also add your 
developer machine’s IP to the “*allow list*”, if you want to have a more 
production-like behavior.