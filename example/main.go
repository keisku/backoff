package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/kskumgk63/backoff"
)

func alwaysErr() error {
	return errors.New("internal server error")
}

func main() {
	cmd := backoff.NewCommander(
		backoff.Timeout(10*time.Second),
		backoff.DebugModeOn(),
		backoff.DebugPrint(func(err error) {
			fmt.Printf("[ERR] %+v\n", err)
		}),
	)
	if timeoutErr := cmd.Exec(alwaysErr); timeoutErr != nil {
		fmt.Println(timeoutErr)
	}
}
