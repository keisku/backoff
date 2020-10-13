package backoff

import (
	"time"
)

type options struct {
	timeout     time.Duration
	maxWaitTime time.Duration
	debugMode   bool
	debugPrint  func(error)
	ignoreError func(error) bool
	timePrint   func(time.Duration)
}

type option func(*options)

// Timeout is to exit the backoff
func Timeout(t time.Duration) option {
	return func(o *options) {
		o.timeout = t
	}
}

// MaxWaitTime sets the maximum waiting time for the backoff interval
func MaxWaitTime(t time.Duration) option {
	return func(o *options) {
		o.maxWaitTime = t
	}
}

func DebugModeOn() option {
	return func(o *options) {
		o.debugMode = true
	}
}

func DebugPrint(f func(error)) option {
	return func(o *options) {
		o.debugPrint = f
	}
}

func IgnoreError(f func(error) bool) option {
	return func(o *options) {
		o.ignoreError = f
	}
}

// TimePrint prints sleepting time
// default: "waiting ?s..."
func TimePrint(f func(time.Duration)) option {
	return func(o *options) {
		o.timePrint = f
	}
}
