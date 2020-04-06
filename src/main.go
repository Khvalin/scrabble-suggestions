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
	"github.com/Khvalin/scrabble-suggestions/src/patterns"
	"github.com/Khvalin/scrabble-suggestions/src/suggestions"
)

type settings struct {
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

	suggestions.LoadDict(string(dictData))

	b := readBoard(settings.Board)
	variants := board.GetVariants(b)
	regs := patterns.ConvertVariantsToRegexes(variants, utf8.RuneCountInString(settings.Hand))

	for _, p := range regs {
		r := suggestions.Match(settings.Hand, p)
		if len(r) > 0 {
			fmt.Println(p, r)
		}
	}
}
