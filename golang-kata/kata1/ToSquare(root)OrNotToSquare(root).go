package kata1

import "math"

func Kata1(input []int) []int {
	result := []int{}

	for _, e := range input {
		if e == 2 {
			result = append(result, e*e)
		} else {
			result = append(result, int(math.Sqrt(float64(e))))
		}
	}
	return result
}
