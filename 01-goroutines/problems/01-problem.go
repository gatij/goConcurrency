/*
Title: "Hello Many Worlds"
Task: Launch 10 goroutines, each printing its own number (1..10). Show how output demonstrates concurrency.
*/

package main

import (
	"fmt"
	"time"
)

func helloManyWorld(n int) {
	fmt.Printf("Hello World from Goroutine: %d\n", n)
}

func main() {
	for i := 1; i <= 10; i++ {
		go helloManyWorld(i)
	}
	time.Sleep(1 * time.Second)
}


// OUTPUT: 
// Each goroutine prints its number, but order is unpredictable due to concurrency.
/*
➜  goConcurrency go run 01-goroutines/problems/01-problem.go  
Hello World from Goroutine: 3
Hello World from Goroutine: 1
Hello World from Goroutine: 2
Hello World from Goroutine: 5
Hello World from Goroutine: 4
Hello World from Goroutine: 6
Hello World from Goroutine: 8
Hello World from Goroutine: 7
Hello World from Goroutine: 9
Hello World from Goroutine: 10
*/
