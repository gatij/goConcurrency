/*
The below code demonstrates the use of WaitGroups to synchronize two goroutines (workers).
Each worker prints a message when it finishes, and the main function waits for both to 
complete before printing a final message.

This ensures that the main program does not exit before the goroutines finish their execution.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Worker 1 finished")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Worker 2 finished")
	}()

	wg.Wait()
	fmt.Println("All workers completed")
}

// Output:

/*
➜  goConcurrency git:(main) go run 02-waitgroups/01-two-workers.go 
Worker 2 finished
Worker 1 finished
All workers completed
*/