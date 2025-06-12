//go:build !solution

package speller

import (
	"fmt"
	"math"
	"strings"
)

var (
	scales = [...]string{"billion", "million", "thousand", ""}
	smalls = [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens   = [...]string{"zero", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
)

func spellToThousand(n int) string {
	var b strings.Builder

	if n >= 100 {
		fmt.Fprintf(&b, "%v hundred", smalls[n/100])
	}

	n = n % 100

	if n > 0 || b.Len() == 0 {
		if b.Len() > 0 {
			b.WriteString(" ")
		}

		if n < len(smalls) {
			b.WriteString(smalls[n])
		} else {
			b.WriteString(tens[n/10])
			if n%10 != 0 {
				fmt.Fprintf(&b, "-%v", smalls[n%10])
			}
		}
	}

	return b.String()
}

func Spell(n int64) string {
	var b strings.Builder
	if n < 0 {
		b.WriteString("minus")
		n = -n
	}

	for i, suffix := range scales {
		pow := (len(scales) - i - 1) * 3
		threshold := int64(math.Pow10(pow))

		if threshold == 1 {
			threshold = 0
		}

		if n >= threshold {
			toThousand := int(n % 1000)
			if threshold > 0 {
				toThousand = int((n / threshold) % 1000)
			}

			if toThousand > 0 || b.Len() == 0 {
				if b.Len() > 0 {
					b.WriteString(" ")
				}
				b.WriteString(spellToThousand(toThousand))

				if len(suffix) > 0 {
					b.WriteString(" ")
					b.WriteString(suffix)
				}
			}
		}
	}

	return b.String()
}
