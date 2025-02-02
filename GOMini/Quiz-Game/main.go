package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./QA.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
