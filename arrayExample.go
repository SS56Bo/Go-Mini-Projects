package main

import "fmt"

func main() {
	var a [10]int
	fmt.Println(a)

	a[5] = 150
	fmt.Println("get:", a)
	fmt.Println("set: ", a[5])
	a[2] = 62
	fmt.Println("get:", a)
	fmt.Println("set: ", a[2])
	fmt.Println("len", len(a))

	b := [5]int{100, 500, 800, 120, 40}
	fmt.Println(b)
	ob := [...]int{1, 7, 5, 4, 20, 17, 24, 84}
	fmt.Println(ob)
	fmt.Println(len(ob))

	// 2D
	var fr [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fr[i][j] = i + j
		}
	}
	fmt.Println(fr)
}
