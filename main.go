package main

import "fmt"

func minus(a int, b int) int {
	ans := a - b
	fmt.Printf("a - b = %d - %d = %d", a, b, ans)
	return ans
}

func main() {
	fmt.Println("hello develop")
}
