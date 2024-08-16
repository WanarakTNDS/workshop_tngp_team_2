package main

import "fmt"

func main() {
	fmt.Println("hello develop")

	fmt.Println("a + b = 40 + 50 = ", Sum(40, 50))
}

func Sum(a, b int) int {
	return a + b
}
