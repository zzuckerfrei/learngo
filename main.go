package main

import (
	"fmt"

	"github.com/zzuckerfrei/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "1"}
	definition, error := dictionary.Search("first")
	if error == nil {
		fmt.Println(definition)
	} else {
		fmt.Println(error)
	}

	word := "hello"
	definition = "my friend"
	error = dictionary.Add(word, definition)
	if error != nil {
		fmt.Println(error)
	}
	value, _ := dictionary.Search(word)
	fmt.Println("found :", word, "definition : ", value)
}
