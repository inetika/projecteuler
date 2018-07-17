package main

import (
	"fmt"
)

func main() {
	card := map[int]int{0: 0, 1: 3, 2: 3, 3: 5, 4: 4, 5: 4, 6: 3, 7: 5, 8: 5, 9: 4, 10: 3, 11: 6, 12: 6, 13: 8, 14: 8, 15: 7, 16: 7, 17: 9, 18: 8, 19: 8, 20: 6, 30: 6, 40: 5, 50: 5, 60: 5, 70: 7, 80: 6, 90: 6, 100: 7, 1000: 8}
	total := 0
	for i := 1; i <= 1000; i++ {
		if i == 1000 {
			total += card[1] + card[1000]
		} else if i >= 100 {
			total += card[i/100] + card[100]
			rest := i % 100
			if rest > 0 {
				total += 3
			}
			if rest <= 20 {
				total += card[rest]
			} else {
				total += card[(rest%100)/10*10] + card[rest%10]
			}
		} else if i > 20 {
			total += card[i/10*10] + card[i%10]
		} else {
			total += card[i]
		}
	}
	fmt.Println(total)
}

/*
Number letter counts
Problem 17
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen)
contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
*/
