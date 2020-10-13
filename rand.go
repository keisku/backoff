package backoff

import (
	"math/rand"
	"time"
)

func randomMilliSecond() time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration(rand.Int63n(1000)) * time.Millisecond
}
