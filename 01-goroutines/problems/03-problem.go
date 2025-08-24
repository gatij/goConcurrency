/*
Title: "Order? Who Knows?"
Task: Have 5 goroutines each print a single statement with a random sleep between 0-500ms before printing. 
      Do you ever get the same output order twice? Why?
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printWithDelay(id int) {
	sleepDuration := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(sleepDuration)
	fmt.Printf("Goroutine %d finished after %v\n", id, sleepDuration)
}

func main() {
	for i := 1; i <= 5; i++ {
		go printWithDelay(i)
	}
	time.Sleep(1 * time.Second) // Wait for all goroutines to finish
}

// OUTPUT: 
// Each run produces a different order due to random sleep durations and concurrency.
/*
➜  goConcurrency go run 01-goroutines/problems/03-problem.go
Goroutine 1 finished after 141ms
Goroutine 5 finished after 253ms
Goroutine 4 finished after 324ms
Goroutine 2 finished after 428ms
Goroutine 3 finished after 434ms
➜  goConcurrency go run 01-goroutines/problems/03-problem.go
Goroutine 4 finished after 61ms
Goroutine 3 finished after 207ms
Goroutine 2 finished after 221ms
Goroutine 5 finished after 249ms
Goroutine 1 finished after 364ms
*/

