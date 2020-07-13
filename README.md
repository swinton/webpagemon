# `webpagemon`
> :eyes: monitor a web page for changes

## Usage

### Compilation

I run this on a Rasperry Pi, to compile for Raspian:

```shell
env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/webpagemon main.go
```

### Wrapper script

I use a wrapper script (mine is named `is-it-in-stock-yet.sh`) to set the environment variables / arguments I need and to call `webpagemon`:

#### `is-it-in-stock-yet.sh`:

```bash
#!/bin/bash
set -e

export TWILIO_ACCOUNT_SID=# your Twilio account SID from, www.twilio.com/console
export TWILIO_AUTH_TOKEN=# your Twilio auth Token from, www.twilio.com/console
export TWILIO_SMS_SENDER=# your Twilio phone number, from www.twilio.com/console/phone-numbers
export WPM_URL=# the web page URL to monitor, e.g. https://www.example.com/
export WPM_SELECTOR=# the document query selector to use to detect a change, e.g. "h1"
export WPM_RECIPIENTS=# comma-separated list of recipient phone numbers, who will be sent a text message when webpagemon detects a change, e.g. `+12223334444,+15556667777`

/path/to/webpagemon \
        "${WPM_URL}" \
        "${WPM_SELECTOR}" \
        "${WPM_RECIPIENTS}"
```


### Crontab

I use cron to periodically run my wrapper script (every 5 minutes here):

```
# m h  dom mon dow   command
*/5 *  *  *    *     /home/pi/bin/is-it-in-stock-yet.sh > /home/pi/logs/webpagemon.log 2>&1
```
