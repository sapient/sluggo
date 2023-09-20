package sluggo

import (
	"unicode"
)

type Options struct {
	Lowercase       bool
	Trim            bool
	ReplacementChar rune
}

type OptionSet func(*Options)

func defaultOptions() Options {
	return Options{
		Lowercase:       true,
		Trim:            true,
		ReplacementChar: '-',
	}
}

func Sluggo(frag string, opts ...OptionSet) string {
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

	if i := len(runes) - 1; runes[i] == '-' {
		runes = runes[:i]
	}

	return string(runes)
}
