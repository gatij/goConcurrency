/*
Title: "Parallel Squares"
Task: Write a function that squares an integer and prints the result. 
      Use goroutines to compute squares of 1..5 in parallel.
*/

package main

import (
	"fmt"
	"time"
)

func square(n int, start time.Time) {
	time.Sleep(1 * time.Second) // Simulate a time-consuming task
	fmt.Printf("Square of %d is %d\n", n, n*n)
	fmt.Println("Time taken for", n, ":", time.Since(start))
}

func main() {

	now := time.Now()
	// Sequential execution
	for i :=1; i <= 5; i++ {
		square(i,now)
	}
	
	fmt.Println("----- Now with Goroutines -----")
	now = time.Now()
	// Parallel execution with goroutines
	for i := 1; i <= 5; i++ {
		go func(num int, start time.Time) {
			square(num, start)
		}(i, now)
	}
	// Wait to allow all goroutines to finish
	time.Sleep(5 * time.Second)
}


// OUTPUT:
// Sequential execution takes ~5 seconds, while parallel execution with goroutines takes ~1 second.
/*
➜  goConcurrency go run 01-goroutines/problems/04-problem.go
Square of 1 is 1
Time taken for 1 : 1.000494625s
Square of 2 is 4
Time taken for 2 : 2.001338125s
Square of 3 is 9
Time taken for 3 : 3.002388917s
Square of 4 is 16
Time taken for 4 : 4.003694208s
Square of 5 is 25
Time taken for 5 : 5.0048615s
----- Now with Goroutines -----
Square of 3 is 9
Square of 4 is 16
Time taken for 3 : 1.000736291s
Square of 2 is 4
Time taken for 2 : 1.000795208s
Square of 1 is 1
Time taken for 1 : 1.000814666s
Time taken for 4 : 1.000762541s
Square of 5 is 25
Time taken for 5 : 1.000898583s
*/