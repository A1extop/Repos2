package main

import (
	"fmt"
)

func Sum(a int, b int) int {
	return a + b
}

func Divide(a int, b int) int {
	return a / b
}
func main() {
	fmt.Println(Sum(2, 3))
}
