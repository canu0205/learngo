package main

import (
	"fmt"

	"github.com/canu0205/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	baseword := "hello"
	definition := "Greeting word"

	if err := dictionary.Add(baseword, definition); err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseword)
	fmt.Println("Result of Add", word)

	if err := dictionary.Delete(baseword); err != nil {
		fmt.Println(err)
	}
	word2, _ := dictionary.Search(baseword)
	fmt.Println("Result of Delete", word2)
}
