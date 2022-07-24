package fibonacci
func Fibonacci(n uint32) uint64 {
	result := [2]uint64{0,1}
	if n <= 1 {
		return 1
	}
	for i := uint32(2); i <= n; i++ {
		result[i % 2] = result[(i - 2)%2] + result[(i - 1)%2]
	}
	return result[n % 2]
}