// Write a function multiplier(factor int) func(int) int that returns a closure.
// The closure should multiply its argument by the factor.
package main

import "fmt"

func multiple(n int) func(mult int) int {
	store := n
	return func(mult int) int {
		return mult * store
	}
}

func main() {
	var Num int
	_, err := fmt.Scan(&Num)
	if err != nil {
		panic("Invalid Input")
	}
	double := multiple(Num)
	fmt.Println(double(5))
	fmt.Println(double(17))
}
