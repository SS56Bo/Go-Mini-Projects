package main

import "fmt"

func fibo(N int) int {
	if N < 2 {
		return N
	}
	return fibo(N-1) + fibo(N-2)
}

func main() {
	V := 6
	fmt.Println(fibo(V))
}
