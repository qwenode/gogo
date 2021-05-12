package rrand

//prevents forgetting to seed
import (
	math_rand "math/rand"
)

// Intn returns, as an int, a non-negative pseudo-random number in [0,n)
// from the default Source.
// It panics if n <= 0.
func Intn(n int) int {
	return math_rand.Intn(n)
}
func Int63() int64 {
	return math_rand.Int63()
}
func Int() int {
	return math_rand.Int()
}
func Int31() int32 {
	return math_rand.Int31()
}
func Int63n(n int64) int64 {
	return math_rand.Int63n(n)
}
func Int31n(n int32) int32 {
	return math_rand.Int31n(n)
}
func Uint32() uint32 {
	return math_rand.Uint32()
}
func Uint64() uint64 {
	return math_rand.Uint64()
}
func Float64() float64 {
	return math_rand.Float64()
}
func Float32() float32 {
	return math_rand.Float32()
}
func ExpFloat64() float64 {
	return math_rand.ExpFloat64()
}

func NormFloat64() float64 {
	return math_rand.NormFloat64()
}
func Shuffle(n int, swap func(i, j int)) {
	math_rand.Shuffle(n, swap)
}
func Read(p []byte) (n int, err error) {
	return math_rand.Read(p)
}
func Perm(n int) []int {
	return math_rand.Perm(n)
}
