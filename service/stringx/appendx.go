package stringx

import (
	"strings"
)

func Appendx(s *[]string, cs ...string) {
	for _, c := range cs {
		c = strings.TrimSpace(c)
		if !IsEmpty(c) {
			if !IsContainx(s, c) {
				*s = append(*s, c)
			}
		}
	}
}
