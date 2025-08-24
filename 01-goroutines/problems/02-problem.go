/*
Title: "Name Shouter"

Task: Write a function that takes a name and prints a greeting. 
	  Launch 5 goroutines with different names. 
	  Ensure main doesn’t exit early. (synchronization with Sleep is acceptable here)
*/

package main

import (
	"fmt"
	"time"
)

func shoutName(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	names := []string{"Alice", "Bob", "Charlie", "Dave", "Eve"}
	for _, name := range names {
		go shoutName(name)
	}
	time.Sleep(1 * time.Second)
}

// OUTPUT: 
// Each goroutine prints its greeting, but order is unpredictable due to concurrency.

/*
➜  goConcurrency go run 01-goroutines/problems/02-problem.go
Hello, Bob!
Hello, Eve!
Hello, Charlie!
Hello, Dave!
Hello, Alice!
➜  goConcurrency go run 01-goroutines/problems/02-problem.go
Hello, Bob!
Hello, Eve!
Hello, Alice!
Hello, Dave!
Hello, Charlie!
*/