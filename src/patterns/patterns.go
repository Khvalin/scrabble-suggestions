package patterns

import (
	"fmt"
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

		for r > l && !isLetter(v[r]) {
			r--
		}
		if l > r {
			continue
		}

		c := 0
		for i := l + 1; i < r; i++ {
			if !isLetter(v[i]) {
				v[i] = '.'
				c++
			}
		}

		if c > maxSubstitutions {
			continue
		}

		minLen := r - l + 1

		if c == 0 {
			minLen++
		}

		str := strings.ReplaceAll(string(v), "#", ".?")

		res = append(res, fmt.Sprintf("^(%s)$", str))
	}

	return res
}
