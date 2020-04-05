package suggestions

import (
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"
)

const ABC = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

var words []string
var wordMap [][]uint8

func countLetters(w string) ([]uint8, bool) {
	aCode, _ := utf8.DecodeRuneInString(ABC)
	abcLen := utf8.RuneCountInString(ABC)
	r := make([]uint8, abcLen)

	ok := true
	for _, ch := range w {
		ind := int(ch) - int(aCode)
		if ind < 0 || ind >= len(r) {
			ok = false
			continue
		}
		r[ind]++
	}

	return r, ok
}

func LoadDict(data string) {
	dictWords := strings.Split(data, "\n")
	words = make([]string, 0, len(dictWords))
	wordMap = make([][]uint8, 0, len(dictWords))

	for _, w := range dictWords {
		w = strings.ToLower(strings.ReplaceAll(w, "ё", "е"))
		if len(w) > 0 {
			if m, ok := countLetters(w); ok {
				words = append(words, string(w))
				wordMap = append(wordMap, m)
			}
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

	needle, _ := countLetters(letters + pattern)
	wildCartCount := 0
	for _, ch := range letters {
		if ch == '*' {
			wildCartCount++
		}
	}

	for i, m := range wordMap {
		count := 0

		if re != nil && !re.MatchString((words[i])) {
			continue
		}

		for ind, c := range m {
			if c > needle[ind] {
				count += int(c) - int(needle[ind])
			}
		}

		if count <= wildCartCount {
			r = append(r, string(words[i]))
		}
	}

	sort.Slice(r, func(i, j int) bool {
		return len(r[i]) < len(r[j])
	})

	return r
}
