package suggestions

import (
	"regexp"
	"sort"
	"strings"
	"unicode"
)

var words []string
var wordMap map[int]map[rune]int

func LoadDict(data string) {
	wordMap = map[int]map[rune]int{}

	words = strings.Split(data, "\n")
	for i, w := range words {
		words[i] = strings.ReplaceAll(w, "ё", "е")
		w = words[i]
		wordMap[i] = map[rune]int{}

		for _, ch := range w {
			wordMap[i][ch]++
		}
	}

}

// Match func
func Match(letters, pattern string) []string {
	r := []string{}

	var re *regexp.Regexp
	if len(pattern) > 0 {
		re = regexp.MustCompile(pattern)
	}

	needle := map[rune]int{}
	for _, ch := range letters {
		needle[ch]++
	}

	for _, ch := range pattern {
		if unicode.IsLetter(ch) {
			needle[ch]++
		}
	}

	for i, m := range wordMap {
		count := 0

		if re != nil && !re.MatchString(words[i]) {
			continue
		}

		for ch, c := range m {
			if c > needle[ch] {
				count++
			}
		}

		if count <= needle['*'] {
			r = append(r, words[i])
		}
	}

	sort.Slice(r, func(i, j int) bool {
		return len(r[i]) < len(r[j])
	})

	return r

}
