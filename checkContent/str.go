package checkContent

import "strings"

func IsEmptyString(str *string) bool {
	if str == nil {
		return true
	}
	return len(strings.TrimSpace(*str)) == 0
}
