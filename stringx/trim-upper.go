package stringx

import "strings"

func TrimUpper(s string) string {
	return strings.ToUpper(strings.TrimSpace(s))
}
