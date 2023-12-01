package kata2

func Divisors(n int) int {
	canDivisorOfInput := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			canDivisorOfInput = append(canDivisorOfInput, i)
		}
	}
	return len(canDivisorOfInput)
}
