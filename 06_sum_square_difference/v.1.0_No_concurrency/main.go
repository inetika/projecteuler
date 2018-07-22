package main

import (
	"fmt"
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	var sumSquare int
	var squareSum int
	defer timeTrack(time.Now(), "10000000000")
	for i := 1; i <= 10000000000; i++ {
		sumSquare += i * i
	}

	for i := 1; i <= 10000000000; i++ {
		squareSum += i
	}
	squareSum = squareSum * squareSum

	fmt.Println(squareSum - sumSquare)
}
