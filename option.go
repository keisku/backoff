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

// Option enables customize the backoff
type Option func(*options)

// Timeout is to exit the backoff
func Timeout(t time.Duration) Option {
	return func(o *options) {
		o.timeout = t
	}
}

// MaxWaitTime sets the maximum waiting time for the backoff interval
func MaxWaitTime(t time.Duration) Option {
	return func(o *options) {
		o.maxWaitTime = t
	}
}

// DebugModeOn enables to print logs
func DebugModeOn() Option {
	return func(o *options) {
		o.debugMode = true
	}
}

func DebugPrint(f func(error)) Option {
	return func(o *options) {
		o.debugPrint = f
	}
}

func IgnoreError(f func(error) bool) Option {
	return func(o *options) {
		o.ignoreError = f
	}
}

// TimePrint prints sleepting time
// default: "waiting ?s..."
func TimePrint(f func(time.Duration)) Option {
	return func(o *options) {
		o.timePrint = f
	}
}
