package backoff

import (
	"time"
)

type options struct {
	timeout           time.Duration
	timeoutErrMessage string
	maxWaitTime       time.Duration
	debugMode         bool
	debugPrint        func(error)
	abortLoop         func(error) bool
	timePrint         func(time.Duration)
}

// Option enables customize the backoff
type Option func(*options)

// Timeout is to exit the backoff
// default: 65 seconds
func Timeout(t time.Duration) Option {
	return func(o *options) {
		o.timeout = t
	}
}

// TimeoutErrMessage is to be output on timeerror
// default: "A timeout ends the exponential backoff"
func TimeoutErrMessage(msg string) Option {
	return func(o *options) {
		o.timeoutErrMessage = msg
	}
}

// MaxWaitTime sets the maximum waiting time for the backoff interval
// default: 32 seconds
func MaxWaitTime(t time.Duration) Option {
	return func(o *options) {
		o.maxWaitTime = t
	}
}

// DebugModeOn enables to print logs
// default: debug mode is off
func DebugModeOn() Option {
	return func(o *options) {
		o.debugMode = true
	}
}

// DebugPrint customizes how to print error
func DebugPrint(f func(error)) Option {
	return func(o *options) {
		o.debugPrint = f
	}
}

// AbortLoop aborts backoff loop when a specific error occurs
func AbortLoop(f func(error) bool) Option {
	return func(o *options) {
		o.abortLoop = f
	}
}

// TimePrint prints sleepting time
// default: "waiting ?s..."
func TimePrint(f func(time.Duration)) Option {
	return func(o *options) {
		o.timePrint = f
	}
}

// DisabelTimePrint disables printing sleep time
func DisabelTimePrint() Option {
	return func(o *options) {
		o.timePrint = func(t time.Duration) {}
	}
}
