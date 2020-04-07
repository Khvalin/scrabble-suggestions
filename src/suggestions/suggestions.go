package suggestions

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/Khvalin/scrabble-suggestions/src/types"
)

const ABC = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

var words []string
var wordMap [][]uint8

func countLetters(w string) ([]uint8, bool) {
	aCode, _ := utf8.DecodeRuneInString(ABC)
	abcLen := utf8.RuneCountInString(ABC)
	r := make([]uint8, abcLen)

	ok := true
	for _, ch := range []rune(w) {
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
func Match(letters, pattern string) []types.MatchResult {
	var res []types.MatchResult

	var re *regexp.Regexp
	if len(pattern) > 0 {
		re = regexp.MustCompile(pattern)
	}

	needle, _ := countLetters(letters + pattern)
	patternMap, _ := countLetters(pattern)
	wildCartCount := 0
	for _, ch := range letters {
		if ch == '*' {
			wildCartCount++
		}
	}

	for i, m := range wordMap {
		var subsCount uint8
		count := 0

		for ind, c := range m {
			if c > needle[ind] {
				count += int(c) - int(needle[ind])
			} else {
				subsCount += c - patternMap[ind]
			}
		}

		if count > wildCartCount || subsCount == 0 {
			continue
		}

		if re == nil || re.MatchString(words[i]) {
			res = append(res, types.MatchResult{Word: words[i], SubtitutionsCount: int(subsCount)})
		}
	}

	return res
}
