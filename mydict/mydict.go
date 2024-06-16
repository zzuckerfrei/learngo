package mydict

import (
	"errors"
	"fmt"
)

// kind of alias
type Dictionary map[string]string

var (
	errorNotFound     = errors.New("not found")
	errorWordExists   = errors.New("that word already exists")
	errorCannotUpdate = errors.New("can't update")
	errosCannotDelete = errors.New("cannot delete")
)

// method receiver
// Go에서 메소드 수신자(Method Receiver)를 사용하여 특정 타입의 객체에서만 메소드를 호출할 수 있도록 만듭니다.
// 메소드 수신자는 메소드를 특정 타입에 바인딩하며, 그 타입의 인스턴스(객체)에서만 해당 메소드를 호출할 수 있게 합니다

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

// Update a word
func (d Dictionary) Update(word, def string) error {
	_, error := d.Search(word)
	switch error {
	case nil:
		d[word] = def
	case errorNotFound:
		return errorCannotUpdate
	}
	return nil
}

// Delete a word
// delete(object, "key")
// the delete function doesn't return anything, and will do nothing if the key doesn't exists
func (d Dictionary) Delete(word string) error {
	_, error := d.Search(word)
	if error == errorNotFound {
		return errosCannotDelete
	}
	delete(d, word)
	return nil
}

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
