package stringx

import (
	"fmt"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Money(v float64) string {

	// v = 1000000.23456
	p := message.NewPrinter(language.English)

	// 1,000,000.23456
	s := p.Sprintf("%f", v)

	items := strings.Split(s, `.`)
	if len(items) > 1 {
		d := strings.TrimRight(items[1], `0`)
		if d != `` {
			// 1,000,000.23000 => 1,000,000.23
			return fmt.Sprintf(`%s.%s`, items[0], d)
		}

		// 1,000,000.00000 => 1,000,000
		return items[0]
	}

	return s

}
