package sluggo

import (
	"unicode"
)

func Sluggo(frag string) string {
	runes := make([]rune, 0, len(frag))
	dashed := false

	for i, c := range frag {
		switch {
		case c >= 'A' && c <= 'Z':
			runes = append(runes, unicode.ToLower(c))
			dashed = false
		case c >= 'a' && c <= 'z', c >= '0' && c <= '9':
			runes = append(runes, c)
			dashed = false
		default:
			if (i < len(frag)-1) && !dashed {
				runes = append(runes, '-')
				dashed = true
			}
		}
	}

	return string(runes)
}
