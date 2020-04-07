package output

import (
	"sort"
	"fmt"

	"github.com/Khvalin/scrabble-suggestions/src/types"
)

func PrintMatchResultsToConsole(matches []types.MatchResult) {
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].SubtitutionsCount < matches[j].SubtitutionsCount
	})

	fmt.Printf("%v\n", matches)
}