package main

import "fmt"

func mod(a, b int) int {
	result := a % b
	return result
}

func main() {
	fmt.Println("hello develop")

	fmt.Printf("a + b = 40 + 50 = %d", mod(40, 50))
}
