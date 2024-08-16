package main

import "fmt"

func main() {
	fmt.Println("hello develop")

	fmt.Printf("a * b = 40 * 50 = %d\n", multiply(40, 50))
}

func multiply(a int, b int) int {
	return a * b
}
