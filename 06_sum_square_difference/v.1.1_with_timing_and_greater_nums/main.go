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
	func() {
		defer timeTrack(time.Now(), "100")
		for i := 1; i <= 100; i++ {
			sumSquare += i * i
		}

		for i := 1; i <= 100; i++ {
			squareSum += i
		}
		squareSum = squareSum * squareSum

		fmt.Println(squareSum - sumSquare)
	}()

	sumSquare, squareSum = 0, 0

	func() {
		defer timeTrack(time.Now(), "100.2")
		for i := 1; i <= 100; i++ {
			sumSquare += i * i
		}

		for i := 1; i <= 100; i++ {
			squareSum += i
		}
		squareSum = squareSum * squareSum

		fmt.Println(squareSum - sumSquare)
	}()

	sumSquare, squareSum = 0, 0

	func() {
		defer timeTrack(time.Now(), "100.3")
		for i := 1; i <= 100; i++ {
			sumSquare += i * i
		}

		for i := 1; i <= 100; i++ {
			squareSum += i
		}
		squareSum = squareSum * squareSum

		fmt.Println(squareSum - sumSquare)
	}()

	sumSquare, squareSum = 0, 0

	func() {
		defer timeTrack(time.Now(), "1,000,000,000")
		for i := 1; i <= 1000000000; i++ {
			sumSquare += i * i
		}

		for i := 1; i <= 1000000000; i++ {
			squareSum += i
		}
		squareSum = squareSum * squareSum

		fmt.Println(squareSum - sumSquare)
	}()

	sumSquare, squareSum = 0, 0

	func() {
		defer timeTrack(time.Now(), "10,000,000,000")
		for i := 1; i <= 10000000000; i++ {
			sumSquare += i * i
		}

		for i := 1; i <= 10000000000; i++ {
			squareSum += i
		}
		squareSum = squareSum * squareSum

		fmt.Println(squareSum - sumSquare)
	}()
}
