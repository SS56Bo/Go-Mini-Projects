package main

import "fmt"

const PI = 3.142

func calculateArea(l int, b int) int {
	return l * b
}

func calculateAreaCircle(radius float32) float32 {
	return PI * radius * radius
}

func calculatePerimeter(l int, b int) int {
	return 2 * (l + b)
}

func calculateCircum(radius float32) float32 {
	return 2 * PI * radius
}

func main() {
	result := calculatePerimeter(120, 45)
	fmt.Println("The perimeter is:", result)
	result = calculateArea(17, 48)
	fmt.Println("The area is:", result)

	var rad float32
	_, err := fmt.Scan(&rad)
	if err != nil {
		panic("Invalid Input")
	}
	resultCirc := calculateAreaCircle(rad)
	fmt.Println("The area is:", resultCirc)
	resultCirc = calculateCircum(rad)
	fmt.Println("The circumference is:", resultCirc)
}
