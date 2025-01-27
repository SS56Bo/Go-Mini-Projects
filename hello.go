package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	if num%5 == 0 {
		fmt.Println("Fizz")
	} else if num%3 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println("FizzBuzz")
	}
}
