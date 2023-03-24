package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1) // allocate only one logical processor

	var wg sync.WaitGroup
	wg.Add(2) // add 2 to wait for 2 goroutines

	fmt.Println("Starting go routines")
	// creating an anoymous function that print the alphabets
	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\n Terminating program...")
}
