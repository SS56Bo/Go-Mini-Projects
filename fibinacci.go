// Write a function fibonacciGenerator that returns a closure.
// The closure should generate the next number in the Fibonacci sequence each time it is called.
package main

import "fmt"

func fiboGenerator(N int) func() int {
	a, b := 0, 1

	return func() int {
		c := a
		a, b = b, a+b
		return c
	}

}

func main() {
	var N int
	fmt.Print("Enter limit of Fibonacci: ")
	_, err := fmt.Scan(&N)
	if err != nil {
		panic("Invalid Input")
	}
	fib := fiboGenerator(N)
	for i := 0; i < N; i++ {
		fmt.Print(fib(), " ")
	}
}
