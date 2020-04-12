package types

type Variant struct {
	BoardLine []rune
	Pattern   []rune
	Matches   []MatchResult
}

type MatchResult struct {
	Word              string
	SubtitutionsCount int
	Offset            int
}
