package board

import "unicode"

func GetVariants(b [][]rune) []string {
	r := []string{}

	for _, row := range b {
		p := 0
		for j := 0; j < len(row); j++ {

			for j < len(row) && unicode.IsLetter(row[j]) {
				j++
			}

			r = append(r, string(row[p:j+1]))
			p = j + 1
		}
	}

	return r
}
