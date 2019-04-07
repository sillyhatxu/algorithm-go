package shuffle

import (
	"math/rand"
)

//shuffle algorithm -> Knuth-Durstenfeld Shuffle
func Start(array []int) []int {
	for i := len(array) - 1; i >= 0; i-- {
		p := randInt64(0, int64(i))
		a := array[i]
		array[i] = array[p]
		array[p] = a
	}
	return array
}

func randInt64(min, max int64) int64 {
	if min >= max || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}
