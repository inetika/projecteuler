package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	var wg sync.WaitGroup
	var sqSum int
	var smSq int

	defer timeTrack(time.Now(), "10000000000")
	wg.Add(2)
	go func() {
		for i := 1; i <= 10000000000; i++ {
			sqSum += i * i
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i <= 10000000000; i++ {
			smSq += i
		}
		smSq = smSq * smSq
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(smSq - sqSum)
}
