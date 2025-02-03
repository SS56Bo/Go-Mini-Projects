package main

import "fmt"

func isPalindrome(N int) bool {
	store := N
	count := 0
	for N != 0 {
		i := N % 10
		count = (count * 10) + i
		N = N / 10
	}
	if count == store {
		return true
	} else {
		return false
	}
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		panic("Invalid input")
	}
	if isPalindrome(num) {
		fmt.Printf("%d is a palindrome!\n", num)
	} else {
		fmt.Printf("%d is not a palindrome!\n", num)
	}
}
