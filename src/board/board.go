package board

import (
	"unicode"

	"github.com/Khvalin/scrabble-suggestions/src/types"
)

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch)
	//return ch != '#'
}

func wordMaskToVariants(borders [][2]int, line []rune) []types.Variant {
	res := make([]types.Variant, 0, len(borders)>>1)
	prev := -1
	softBorders := make([][2]int, len(borders))
	for i, pair := range borders {
		softBorders[i] = pair
		// left part
		if pair[0] > prev+1 {
			softBorders[i][0] = prev + 1
		}

		r := len(line)
		if i < len(borders)-1 {
			r = borders[i+1][0] - 1
		}

		//right part
		if r > pair[1] {
			softBorders[i][1] = r
		}

		prev = pair[1]
	}

	for i, p := range softBorders {
		res = append(res, types.Variant{BoardLine: append([]rune{}, line[p[0]:p[1]]...)})

		for j := i + 1; j < len(softBorders); j++ {
			res = append(res, types.Variant{BoardLine: append([]rune{}, line[p[0]:softBorders[j][1]]...)})
		}
	}

	return res
}

func GetVariants(b [][]rune) []types.Variant {
	r := []types.Variant{}

	//horizontal
	for _, row := range b {
		c := -1
		var wordBorders [][2]int

		for i, ch := range row {
			if isLetter(ch) {
				if i > 0 && isLetter(row[i-1]) {
					wordBorders[c][1] = i + 1
					continue
				}
				c++
				wordBorders = append(wordBorders, [2]int{i, i + 1})
			}
		}

		r = append(r, wordMaskToVariants(wordBorders, row)...)
	}

	//vertical
	for col := range b[0] {
		c := -1
		var wordBorders [][2]int
		column := make([]rune, len(b[0]))

		for i := range b {
			ch := b[i][col]
			column[i] = ch
			if isLetter(ch) {
				if i > 0 && isLetter(column[i-1]) {
					wordBorders[c][1] = i + 1
					continue
				}

				c++
				wordBorders = append(wordBorders, [2]int{i, i + 1})
			}
		}

		r = append(r, wordMaskToVariants(wordBorders, column)...)
	}

	return r
}
