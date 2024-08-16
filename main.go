package main

import "fmt"

func main() {
	fmt.Println("hello world")
	a, b := 10, 2
	output := devide(a, b)
	fmt.Printf("a / b = %v / %v = %v", a, b, output)
}

func devide(a, b int) int {
	result := a / b
	return result
}
