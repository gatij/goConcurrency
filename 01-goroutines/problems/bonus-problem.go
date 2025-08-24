// Lets simulate a baking of cake using goroutines
// A cake requires many tasks that can be done concurrently while preparing it.
// after concurrent tasks are done, some tasks need to be done sequentially.
// then the cake is baked.

package main

import (
	"fmt"
	"time"
)

// Simulate a task that takes time
func task(name string, duration time.Duration) {
	fmt.Printf("Starting %s...\n", name)
	time.Sleep(duration)
	fmt.Printf("%s done!\n", name)
}

func main() {

	// Concurrent tasks
	go task("Mixing Ingredients", 2*time.Second)
	go task("Preparing Frosting", 1*time.Second)
	go task("Chopping Nuts", 1*time.Second)

	// Wait for concurrent tasks to finish
	time.Sleep(3 * time.Second)

	// Sequential tasks
	task("Assembling Cake", 2*time.Second)
	task("Baking Cake", 3*time.Second)

	fmt.Println("Cake is ready!")
}


// OUTPUT:
// The concurrent tasks (Mixing Ingredients, Preparing Frosting, Chopping Nuts) will complete first, followed by the sequential tasks (Assembling Cake, Baking Cake).
/*
➜  goConcurrency git:(main) ✗ go run 01-goroutines/problems/bonus-problem.go
Starting Mixing Ingredients...
Starting Preparing Frosting...
Starting Chopping Nuts...
Chopping Nuts done!
Preparing Frosting done!
Mixing Ingredients done!
Starting Assembling Cake...
Assembling Cake done!
Starting Baking Cake...
Baking Cake done!
Cake is ready!
*/


// Note: The order of completion for the concurrent tasks may vary with each run, demonstrating concurrency.