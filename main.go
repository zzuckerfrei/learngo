package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

// naked return
// defer : after func execute
func lenAndUpper(name string) (length int, uppper string) {
	length = len(name)
	uppper = strings.ToUpper(name)
	defer fmt.Println("I'm done!\n" + uppper)
	return
}

// for
func supperAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		// fmt.Println(idx, number)
		total += number
	}

	return total
}

// if + variable expression
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // koreanAge is variable expression
		defer fmt.Println("[canIDrink] : you are not allowed to drink :(")
		return false
	}
	return true
}

// switch
// if~else 여러번 사용하는 것 방지
// variable expression 가능
func canIDrinkSwitch(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 20:
		fmt.Println("[canIDrinkSwitch] : not allowed :(")
		return false
	case koreanAge >= 20:
		fmt.Println("[canIDrinkSwitch] : enjoy :)")
		return true
	}
	return false
}

// pointer (low level programming)
// & : memory addr copy
// * : check the value stored at this address
// if you want to handle heavy objects, but you don't want to copy it.
// this makes your program faster.
func pointer() {
	a := 2
	b := &a
	a = 5

	fmt.Println("[pointer] :", a, b)
	fmt.Println("[pointer] :", *b)

	// you can change value with address
	*b = 10
	fmt.Println("[pointer] :", a)
}

// Arrays and Slices
// array : limit length
// slice : unlimit length
func arrayAndSlice() {
	// array
	game := [3]string{"pandemic", "legacy", "fun"}
	fmt.Println("[arrayAndSlice] : ", game)

	// slice
	phone := []string{"I", "am", "iPhone", "12", "mini"}
	fmt.Println("[arrayAndSlice] : ", phone)

	// append : if you want to add value to the slice
	// 'append' doesn't push value. it returns the new slice with the added value.
	phone = append(phone, "blue")
	fmt.Println("[arrayAndSlice] : ", phone)
}

// map
// map[typeOfKey]typeOfvalue{"value":"key"}
func makeMap() {
	zuckerfrei := map[string]int{"age": 99, "eye": 2, "mouth": 1}
	fmt.Println("[makeMap] : ", zuckerfrei)

	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println("[makeMap] : ", nico)

	// iterate map with range
	for _, value := range nico {
		fmt.Println("[makeMap] : ", value)
	}
}

// struct

func main() {
	println("한글도 잘 나오나?")

	fmt.Println(multiply(2, 2))

	repeatMe("test", "test2", "test3")

	// naked return, defer
	name_len, name_upper := lenAndUpper("zuckerfrei")
	fmt.Println(name_len, name_upper)

	// for
	total := supperAdd(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(total)

	// if
	fmt.Println(canIDrink(15))

	// switch
	fmt.Println(canIDrinkSwitch(24))

	// pointer
	pointer()

	// array and slice
	arrayAndSlice()

	// map
	makeMap()
}
