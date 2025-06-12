//go:build !solution

package reverse

import (
	"strings"
	"unicode/utf8"
)

func Reverse(input string) string {
	var b strings.Builder
	b.Grow(len(input))

	for i := len(input); i > 0; {
		r, size := utf8.DecodeLastRuneInString(input[:i])

		b.WriteRune(r)
		i -= size
	}

	return b.String()
}
