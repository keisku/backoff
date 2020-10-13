package backoff

import "math"

func pow2(n int) int {
	return int(math.Pow(2, float64(n)))
}
