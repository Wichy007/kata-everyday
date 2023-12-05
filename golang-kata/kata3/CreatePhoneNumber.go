package kata3

func CreatePhoneNumber(numbers [10]uint) []string {
	result := []string{}
	for i := 0; i < 14; i++ {
		if i == 0 {
			result = append(result, "(")
		} else if i == 4 {
			result = append(result, ")")
		} else if i == 5 {
			result = append(result, " ")
		} else if i == 9 {
			result = append(result, "-")
		} else {
			// result = append(result, strconv.Itoa(numbers[i]))
		}
	}

	return result
}
