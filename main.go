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
// defer
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
}
