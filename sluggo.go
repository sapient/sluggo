package sluggo

import (
	"fmt"
	"unicode"
)

func Sluggo(frag string) string {
	runes := make([]rune, 0, len(frag))

	for i, c := range frag {
		switch {

		case c >= 'A' && c <= 'Z':
			runes = append(runes, unicode.ToLower(c))
		case c >= 'a' && c <= 'z', c >= '0' && c <= '9':
			runes = append(runes, c)
		default:
			fmt.Println("i:", i, "frag:", len(frag))
			if i < len(frag)-1 {
				runes = append(runes, '-')
			}
		}
	}

	return string(runes)
}
