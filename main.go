package main

import (
	"fmt"
	"math"
)

func minus(a int, b int) int {
	ans := a - b
	return ans
}

func mod(a, b int) int {
	result := a % b
	return result
}

func main() {
	fmt.Println("hello develop")

	fmt.Println("a + b = 40 + 50 = ", Sum(40, 50))
	fmt.Println("a - b = 40 - 50 = ", minus(40, 50))
	fmt.Printf("a * b = 40 * 50 = %d\n", multiply(40, 50))
	fmt.Printf("a + b = 40 + 50 = %d\n", mod(40, 50))
	fmt.Printf("a / b = 40 / 50 = %d\n", devide(40, 50))
	fmt.Printf("Sqrt(a) = Sqrt(9) = %.2f\n", squareroot(9))
}

func multiply(a int, b int) int {
	return a * b
}

func Sum(a, b int) int {
	return a + b
}

func devide(a, b int) int {
	if b == 0 {
		panic("can not devide by 0")
	}
	result := a / b
	return result
}

func squareroot(a int) float64 {
	return math.Sqrt(float64(a))
}
