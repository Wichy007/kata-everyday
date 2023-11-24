package kata1

import "math"

func Kata1(input []int) []int {
	result := []int{}

	for _, e := range input {
		if int(math.Sqrt(float64(e)))*int(math.Sqrt(float64(e))) == e {
			result = append(result, int(math.Sqrt(float64(e))))
		} else {
			result = append(result, e*e)
		}
	}
	return result
}
