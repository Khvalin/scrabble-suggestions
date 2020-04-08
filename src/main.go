package main

import (
	// standard library

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/Khvalin/scrabble-suggestions/src/board"
	"github.com/Khvalin/scrabble-suggestions/src/output"
	"github.com/Khvalin/scrabble-suggestions/src/patterns"
	"github.com/Khvalin/scrabble-suggestions/src/suggestions"
)

type settings struct {
	Abc      string
	DictPath string
	Hand     string
	Pattern  string
	Board    []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readBoard(lines []string) [][]rune {
	res := [][]rune{}

	for i, row := range lines {
		res = append(res, make([]rune, utf8.RuneCountInString(row)))

		for j, ch := range []rune(row) {
			res[i][j] = unicode.ToLower(ch)
		}
	}

	return res
}

func readFile(fileName string) []byte {
	dat, e := ioutil.ReadFile(fileName)
	check(e)

	return dat
}

func loadSettings(settingsFileName string) *settings {
	settingsData := readFile(settingsFileName)

	var s settings
	e := json.Unmarshal(settingsData, &s)
	check(e)
	s.Hand = strings.ToLower(s.Hand)

	return &s
}

func main() {
	settings := loadSettings(os.Args[1])
	dictData := readFile(settings.DictPath)

	matcher := suggestions.CreateMatcher(settings.Abc)

	matcher.LoadDict(string(dictData))

	b := readBoard(settings.Board)
	variants := board.GetVariants(b)

	// for _, v := range variants {
	// 	fmt.Println(string(v))
	// }

	regs := patterns.ConvertVariantsToRegexes(variants, utf8.RuneCountInString(settings.Hand))

	if len(variants) == 0 {
		regs = append(regs, "")
	}

	r := matcher.Match(settings.Hand, regs)

	for i, p := range r {
		if len(p) > 0 {
			fmt.Println(regs[i])
			output.PrintMatchResultsToConsole(p)
		}
	}
}
