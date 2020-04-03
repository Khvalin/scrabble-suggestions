package main

import (
	// standard library

	"fmt"
	"io/ioutil"
	"os"

	"github.com/Khvalin/scrabble-suggestions/src/suggestions"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dictFileName := os.Args[1]

	data, _ := ioutil.ReadFile(dictFileName)

	suggestions.LoadDict(string(data))
	r := suggestions.Match("ллнесдр", "^.*о.*$")

	fmt.Println(r)
}
