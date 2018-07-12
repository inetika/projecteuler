package main

import (
	"fmt"
	"time"
)

//main will ask for amount of dividers you need
//and find the smallest multiple while timing itself
//on my machine difference was 41 sec vs 1 sec using 1_simple_fast

func main() {
	var divs int
	fmt.Print("How many dividers? 1 through: ")
	fmt.Scan(&divs)

	t := time.Now()
	fmt.Println(t.Format("20060102150405"))

	i := 1
	dividers := make([]int, divs)
	for key := range dividers {
		dividers[key] = i
		i++
	}

	ranger := func(val int) bool {
		result := 0
		for _, v := range dividers {
			if val%v == 0 {
				result++
			}
		}
		if result == len(dividers) {
			return true
		}
		return false
	}

	for {
		if ranger(i) {
			fmt.Println(i)
			break
		} else {
			i++
			continue
		}
	}
	t = time.Now()
	fmt.Println(t.Format("20060102150405"))
}

// Smallest multiple
// Problem 5
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
