package types

import "fmt"

type MatchResult struct {
	Word              string
	SubtitutionsCount int
}

func (m MatchResult) String() string {
	return fmt.Sprintf("%s %d", m.Word, m.SubtitutionsCount)
}
