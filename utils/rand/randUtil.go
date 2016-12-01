package rand

import "math/rand"


//return rand number between [start, end)
func Rand(start, end int32) int32 {
	if end - start <= 0 {
		return start
	}
	return start + rand.Int31n(end - start)
}

//return rand number between [0, n)
func Randn(n int32) int32 {
	if n <= 0 {
		return 0
	}
	return rand.Int31n(n)
}
