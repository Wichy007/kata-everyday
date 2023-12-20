package kata4

import "strings"

func ToCamelCase(s string) string {
	stringArray := strings.Split(s, "-")
	stringArray[1] = strings.Title(stringArray[1])
	return strings.Join(stringArray, "")
}
