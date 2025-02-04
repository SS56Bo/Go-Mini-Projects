// Write a function counter that returns another function.
// The returned function should increment and return a count each time it is called.
package main

import "fmt"

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	c1 := counter()
	var N int
	fmt.Println("Enter a number: ")
	_, err := fmt.Scan(&N)
	if err != nil {
		panic(err)
	}
	for i := 0; i < N; i++ {
		fmt.Println(c1())
	}
}
