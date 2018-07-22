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

	func() {
		defer timeTrack(time.Now(), "100")
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
	}()

	smSq, sqSum = 0, 0

	func() {
		defer timeTrack(time.Now(), "100.2")
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
	}()

	smSq, sqSum = 0, 0

	func() {
		defer timeTrack(time.Now(), "100.3")
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
	}()

	smSq, sqSum = 0, 0

	func() {
		defer timeTrack(time.Now(), "1,000,000,000")
		wg.Add(2)
		go func() {
			for i := 1; i <= 1000000000; i++ {
				sqSum += i * i
			}
			wg.Done()
		}()

		go func() {
			for i := 1; i <= 1000000000; i++ {
				smSq += i
			}
			smSq = smSq * smSq
			wg.Done()
		}()
		wg.Wait()
		fmt.Println(smSq - sqSum)
	}()

	smSq, sqSum = 0, 0

	func() {
		defer timeTrack(time.Now(), "10,000,000,000")
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
	}()

}
