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
	//Tried switching to map instead of []int
	//Made it slower
	dividers := make([]int, divs)
	for key := range dividers {
		dividers[key] = i
		i++
	}

	/*
	 Although iteration through maps is supposed to be faster
	 here it increases calculation time
	*/
	// dividers := make(map[int]int)
	// for i := 1; i <= divs; i++ {
	// 	dividers[i] = i
	// }

	/*
		V.1.0
		This version of ranger func caused total calculation time to grow to 40 sec
		vs 5 with the current
	*/
	// ranger := func(val int) bool {
	// 	result := 0
	// 	for _, v := range dividers {
	// 		if val%v != 0 {
	// 			result++
	// 		}
	// 	}
	// 	if result == len(dividers) {
	// 		return true
	// 	}
	// 	return false
	// }

	//V.1.1
	ranger := func(val int) bool {
		for _, v := range dividers {
			if val%v != 0 {
				return false
			}
		}
		return true
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
