package backoff

import (
	"time"
)

type options struct {
	timeout      time.Duration
	maxBackoff   time.Duration
	debugMode    bool
	debugPrinter func(error)
	ignoreError  func(error) bool
	timePrinter  func(time.Duration)
}

type option func(*options)

func Timeout(t time.Duration) option {
	return func(o *options) {
		o.timeout = t
	}
}

func MaxBackoff(t time.Duration) option {
	return func(o *options) {
		o.maxBackoff = t
	}
}

func DebugModeOn() option {
	return func(o *options) {
		o.debugMode = true
	}
}

func DebugPrinter(f func(error)) option {
	return func(o *options) {
		o.debugPrinter = f
	}
}

func IgnoreError(f func(error) bool) option {
	return func(o *options) {
		o.ignoreError = f
	}
}

// TimePrinter prints sleepting time
// default: "waiting ?s..."
func TimePrinter(f func(time.Duration)) option {
	return func(o *options) {
		o.timePrinter = f
	}
}
