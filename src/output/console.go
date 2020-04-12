package output

import (
	"fmt"
	"sort"
	"unicode"

	"github.com/Khvalin/scrabble-suggestions/src/patterns"
	"github.com/Khvalin/scrabble-suggestions/src/types"
)

func PrintFinalVariantsToConsole(variants []types.Variant) {
	for _, v := range variants {
		sort.Slice(v.Matches, func(i, j int) bool {
			return v.Matches[i].SubtitutionsCount < v.Matches[j].SubtitutionsCount
		})
	}

	for _, v := range variants {
		if len(v.Matches) == 0 {
			continue
		}

		fmt.Println(string(v.Pattern))
		for _, m := range v.Matches {
			w := []rune(m.Word)
			j := m.Offset

			for i, ch := range w {
				if patterns.IsLetter(v.Pattern[j]) {
					w[i] = unicode.ToUpper(ch)
				}
				j++
			}
			fmt.Printf("{%s %d} ", string(w), m.SubtitutionsCount)
		}
		fmt.Print("\n\n")
	}
}
