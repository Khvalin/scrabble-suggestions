package suggestions

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/Khvalin/scrabble-suggestions/src/types"
)

var words []string
var wordMap [][]uint8

type MatcherInterface interface {
	countLetters(w string) ([]uint8, bool)

	LoadDict(data string)
	Match(letters string, patterns []string) [][]types.MatchResult
}

type Matcher struct {
	abc []rune
}

func CreateMatcher(abc string) MatcherInterface {
	return Matcher{[]rune(abc)}
}

func (matcher Matcher) countLetters(w string) ([]uint8, bool) {
	aCode := matcher.abc[0]
	abcLen := len(matcher.abc)
	r := make([]uint8, abcLen)

	ok := true
	for _, ch := range []rune(w) {
		ind := int(ch) - int(aCode)
		if !unicode.IsLetter(ch) || (ind < 0 || ind >= len(r)) {
			ok = false
			continue
		}
		r[ind]++
	}

	return r, ok
}

func (matcher Matcher) LoadDict(data string) {
	dictWords := strings.Split(data, "\n")
	words = make([]string, 0, len(dictWords))
	wordMap = make([][]uint8, 0, len(dictWords))

	for _, w := range dictWords {
		w = strings.ToLower(strings.ReplaceAll(w, "ё", "е"))
		if len(w) > 0 {
			if m, ok := matcher.countLetters(w); ok {
				words = append(words, string(w))
				wordMap = append(wordMap, m)
			}
		}
	}
}

// Match func
func (matcher Matcher) Match(letters string, patterns []string) [][]types.MatchResult {
	res := make([][]types.MatchResult, len(patterns))

	wildCartCount := 0
	for _, ch := range letters {
		if ch == '*' {
			wildCartCount++
		}
	}

	lettersMap, _ := matcher.countLetters(letters)

	filteredIds := make([]int, 0, len(wordMap))

	for i, m := range wordMap {
		if wildCartCount > 0 {
			filteredIds = append(filteredIds, i)
		} else {
			for j, c := range m {
				if c > 0 && lettersMap[j] > 0 {
					filteredIds = append(filteredIds, i)
					break
				}
			}
		}
	}

	for k, pattern := range patterns {
		var re *regexp.Regexp
		if len(pattern) > 0 {
			re = regexp.MustCompile(pattern)
		}

		needle, _ := matcher.countLetters(letters + pattern)
		patternMap, _ := matcher.countLetters(pattern)

		for _, ind := range filteredIds {
			m := wordMap[ind]

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

			if re == nil || re.MatchString(words[ind]) {
				res[k] = append(res[k], types.MatchResult{Word: words[ind], SubtitutionsCount: int(subsCount)})
			}
		}
	}

	return res
}
