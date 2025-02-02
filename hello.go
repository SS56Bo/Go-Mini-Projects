package main

import "fmt"

// Function to check if a number is prime
func checkPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var N int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	fmt.Println("Prime numbers up to", N, "are:")
	for i := 2; i <= N; i++ {
		if checkPrime(i) {
			fmt.Println(i)
		}
	}
}
