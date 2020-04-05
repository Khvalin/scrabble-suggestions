package main

import (
	// standard library

	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Khvalin/scrabble-suggestions/src/board"
	"github.com/Khvalin/scrabble-suggestions/src/suggestions"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readBoard(data string) [][]rune {
	res := [][]rune{}
	lines := strings.Split(data, "\n")

	for i, row := range lines {
		res = append(res, make([]rune, len(row)))

		for j, ch := range row {

			res[i][j] = ch

		}
	}

	return res
}

func main() {
	dictFileName := os.Args[1]
	boardFileName := os.Args[2]

	dictData, e := ioutil.ReadFile(dictFileName)
	check(e)

	suggestions.LoadDict(string(dictData))
	r := suggestions.Match("ъвлоагз", "^(.)*о(.)*$")
	fmt.Println(r, "\n")

	boardData, e := ioutil.ReadFile(boardFileName)
	check(e)

	b := readBoard(string(boardData))
	variants := board.GetVariants(b)
	fmt.Println(len(variants), variants)
}
