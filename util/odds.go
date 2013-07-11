package util

func Fact(n int64) int64 {
	if n < 0 || n == 0 {
		return 1
	}
	return n * Fact(n-1)
}
