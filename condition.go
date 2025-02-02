package main

import (
	"fmt"
	"reflect"
)

func factorial(N rune) rune {
	if N == 0 {
		return 1
	}
	return N * factorial(N-1)
}

func main() {
	var N rune
	fmt.Println("Enter a number: ")
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println("Invalid Input")
		return
	}
	fmt.Println(factorial(N))
	fmt.Println(reflect.TypeOf(factorial(N)))
}
