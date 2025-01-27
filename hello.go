package main

import "fmt"

func main() {
	num := 89
	if num%5 == 0 {
		fmt.Println("Fizz")
	} else if num%3 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println("FizzBuzz")
	}
}
