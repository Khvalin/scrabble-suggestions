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

func ConvertVariantsToPatterns(variants [][]rune, maxSubstitutions int) [][]rune {
	var res = make([][]rune, 0, len(variants))

	for i := range variants {
		v := append([]rune{}, variants[i]...)

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

		for j := 0; j < len(v); j++ {
			if v[j] == '#' {
				v[j] = '.'
			}
		}

		res = append(res, v)
	}

	return res
}

func MatchPattern(pattern []rune, word string) (bool, int) {
	p := []rune(pattern)
	s := 0
	for s < len(p) && !isLetter(p[s]) {
		s++
	}

	if s == len(p) {
		//TODO
		return false, -1
	}

	e := len(p) - 1
	for e >= 0 && !isLetter(p[e]) {
		e--
	}

	offset := s
	w := []rune(word)

	j, k := 0, 0
outer:
	for i := 0; i < len(w); i++ {
		if w[i] != p[s] {
			continue
		}

		for j, k := i-1, s-1; j >= 0; j-- {
			if k < 0 {
				return false, -1
			}

			if j == 0 {
				offset = k
			}

			//if p[k] == '?' || p[k] == '.' {
			k--
			//}
		}

		for j, k = i+1, s+1; j < len(w); j, k = j+1, k+1 {
			if k >= len(p) {
				break outer
			}
			if isLetter(p[k]) && p[k] != w[j] {
				continue outer
			}
		}
	}

	return k > e && j == len(w), offset
}
