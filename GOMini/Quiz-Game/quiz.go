package main

import (
	"fmt"
	"os"
)

func main() {
	// data, err := os.ReadFile("./QA.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	//can also create a slice and take on each at a time
	f, err := os.Open("./QA.csv")
	if err != nil {
		panic(err)
	}

	b1 := make([]byte, 120)
	n1, err := f.Read(b1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes:\n%s\n", n1, string(b1[:n1]))
}
