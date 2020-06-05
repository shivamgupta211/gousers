package logger

import (
	"gousers/enums"
	"strconv"

	"github.com/getsentry/raven-go"
)

// Scaffold ... COnfiguring the Raven
func Scaffold() {
	sentry := enums.GetSentry()
	raven.SetDSN(sentry)
}

// CaptureError ... Capturing and sent to Sentry
func CaptureError(err error, code int, message string) {
	var tags = make(map[string]string)
	tags["code"] = strconv.Itoa(code)
	tags["message"] = message
	raven.CaptureError(err, tags)
	panic(err)
}
