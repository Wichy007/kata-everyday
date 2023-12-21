package kata4

import (
	"strings"
)

func ToCamelCase(s string) string {
	stringArray := strings.Split(s, "-")
	for i, e := range stringArray {
		if i == 0 {
			continue
		}
		stringArray[i] = strings.Title(e)
	}

	return strings.Join(stringArray, "")
}
