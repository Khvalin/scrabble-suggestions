package patterns

import (
	"strings"
	"unicode"
)

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch)
	//return ch != '#'
}

func ConvertVariantsToRegexes(variants [][]rune, maxSubstitutions int) []string {
	var res = make([]string, 0, len(variants))

	for _, v := range variants {
		l, r := 0, len(v)-1
		for l < len(v)-1 && !isLetter(v[l]) {
			l++
		}

		for r >= 0 && !isLetter(v[r]) {
			r--
		}
		if l > r {
			continue
		}

		c := 0
		for i := l; i <= r; i++ {
			if !isLetter(v[i]) {
				v[i] = '.'
				c++
			}
		}

		if c == 0 && l > 0 && r == len(v)-1 {
			c++
			v[l-1] = '.'
		}

		if c == 0 && l == 0 && r < len(v)-1 {
			c++
			v[r+1] = '.'
		}

		str := strings.ReplaceAll(string(v), "#", "(.?)")

		if c > 0 && c <= maxSubstitutions {
			res = append(res, "^"+str+"$")
		}
	}

	return res
}
