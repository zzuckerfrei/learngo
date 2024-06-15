package mydict

import (
	"errors"
)

// kind of alias
type Dictionary map[string]string

var errorNotFound = errors.New("Not Found")
var errorWordExists = errors.New("That word already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	// fmt.Println(value)
	// fmt.Println(exists)

	if exists {
		return value, nil
	}
	return "", errorNotFound
}

// Add a word to dict
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errorNotFound:
		d[word] = def

	case nil:
		return errorWordExists
	}
	return nil
}
