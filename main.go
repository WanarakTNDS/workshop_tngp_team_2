package main

import "fmt"

func minus(a int, b int) int {
	ans := a - b
	fmt.Printf("a - b = %d - %d = %d", a, b, ans)
	return ans
}

func mod(a, b int) int {
	result := a % b
	return result
}

func main() {
	fmt.Println("hello develop")

	fmt.Println("a + b = 40 + 50 = ", Sum(40, 50))
	fmt.Printf("a * b = 40 * 50 = %d\n", multiply(40, 50))
}

func multiply(a int, b int) int {
	return a * b
}

func Sum(a, b int) int {
	fmt.Printf("a + b = 40 + 50 = %d", mod(40, 50))
	return a + b
}
