package main

import (
	"fmt"
)

func main() {
	var sumSquare int
	var squareSum int

	for i := 1; i <= 100; i++ {
		sumSquare += i * i
	}

	for i := 1; i <= 100; i++ {
		squareSum += i
	}
	squareSum = squareSum * squareSum

	fmt.Println(squareSum - sumSquare)
}
