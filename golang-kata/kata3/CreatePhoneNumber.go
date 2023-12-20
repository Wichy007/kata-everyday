package kata3

import (
	"fmt"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {

	range1 := strings.Join([]string{fmt.Sprintf("%d%d%d", numbers[0], numbers[1], numbers[2])}, "")
	range2 := strings.Join([]string{fmt.Sprintf("%d%d%d", numbers[3], numbers[4], numbers[5])}, "")
	range3 := strings.Join([]string{fmt.Sprintf("%d%d%d%d", numbers[6], numbers[7], numbers[8], numbers[9])}, "")

	return fmt.Sprintf("(%s) %s-%s", range1, range2, range3)
}
