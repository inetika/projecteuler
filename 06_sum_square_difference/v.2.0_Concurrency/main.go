package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var sqSum int
	var smSq int

	wg.Add(2)
	go func() {
		for i := 1; i <= 100; i++ {
			sqSum += i * i
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			smSq += i
		}
		smSq = smSq * smSq
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(smSq - sqSum)
}
