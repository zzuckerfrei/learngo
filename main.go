package main

import (
	"fmt"

	"github.com/zzuckerfrei/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "1"}

	// search
	definition, error := dictionary.Search("first")
	if error == nil {
		fmt.Println(definition)
	} else {
		fmt.Println(error)
	}

	// add
	word := "hello"
	definition = "my friend"
	error = dictionary.Add(word, definition)
	if error != nil {
		fmt.Println(error)
	}
	value, _ := dictionary.Search(word)
	fmt.Println("found :", word, "definition : ", value)

	// update
	word = "first"
	definition = "11"
	error = dictionary.Update(word, definition)
	if error == nil {
		fmt.Println("update key :", word, "value :", dictionary[word])
	} else {
		fmt.Println(error)
	}

	// delete
	word = "first"
	error = dictionary.Delete(word)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(dictionary.Search(word))
}
