package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	fmt.Println("\ncreating goroutines")
	go printPrime("A")
	go printPrime("B")
	fmt.Println("\nwaiting to finish")
	wg.Wait()

	fmt.Println("Terminating")
}

func printPrime(prefix string) {
	// printPrime displays prime numbers for the first 5000 numbers
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Printf("%s completed", prefix)

}
