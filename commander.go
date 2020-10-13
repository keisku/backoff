package backoff

import (
	"errors"
	"fmt"
	"time"
)

type Commander struct{ options }

func NewCommander(opts ...option) Commander {
	var (
		defaultTimeout    = 1 * time.Minute
		defaultmaxBackoff = 32 * time.Second
		defaultDebugMode  = false
		defaultErrPrint   = func(err error) {
			fmt.Println(err)
		}
		defaultIgnoreError = func(error) bool { return false }
		defaultTimePrint   = func(d time.Duration) {
			fmt.Printf("waiting %fs...\n", d.Seconds())
		}
	)
	options := options{
		timeout:     defaultTimeout,
		maxBackoff:  defaultmaxBackoff,
		debugMode:   defaultDebugMode,
		debugPrint:  defaultErrPrint,
		ignoreError: defaultIgnoreError,
		timePrint:   defaultTimePrint,
	}
	for _, opt := range opts {
		opt(&options)
	}
	return Commander{options}
}

func (cmd Commander) Exec(f func() error) error {
	done := make(chan struct{})
	go cmd.backoffLoop(done, f)

	select {
	case <-done:
		return nil
	case <-time.After(cmd.timeout):
		return errors.New("A timeout ends the exponential backoff")
	}
}

func (cmd Commander) backoffLoop(done chan struct{}, f func() error) {
	exponent := 1
	for {
		err := f()
		if err == nil {
			done <- struct{}{}
			return
		}
		if cmd.ignoreError(err) {
			if cmd.debugMode {
				cmd.debugPrint(err)
			}
			done <- struct{}{}
			return
		}
		exponentSecond := time.Duration(pow2(exponent)) * time.Second
		d := min(exponentSecond+randomMilliSecond(), cmd.maxBackoff)
		if cmd.debugMode {
			cmd.debugPrint(err)
			cmd.timePrint(d)
		}
		time.Sleep(d)
		exponent++
	}
}
