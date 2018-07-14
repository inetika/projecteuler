package main

import (
	"fmt"
)

var dateofweek = 1
var sundays = 0

func countSundays(d int, m int, y int) {
	if dateofweek <= 6 {
		dateofweek++
	} else {
		dateofweek = 1
	}
	if d == 1 && dateofweek == 7 && y != 1900 {
		sundays++
	}
}

func main() {

	for y := 1900; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			if m == 4 || m == 6 || m == 9 || m == 11 {
				for d := 1; d <= 30; d++ {
					countSundays(d, m, y)
				}
			} else if m == 2 {
				if y%4 == 0 {
					if y%100 == 0 && y%400 != 0 {
						for d := 1; d <= 28; d++ {
							countSundays(d, m, y)
						}
					} else {
						for d := 1; d <= 29; d++ {
							countSundays(d, m, y)
						}
					}
				} else {
					for d := 1; d <= 28; d++ {
						countSundays(d, m, y)
					}
				}
			} else {
				for d := 1; d <= 31; d++ {
					countSundays(d, m, y)
				}
			}
		}
	}
	fmt.Println(sundays)
}

/*
Counting Sundays
Problem 19
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/
