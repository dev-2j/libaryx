package stringx

import "strings"

func TrimLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
