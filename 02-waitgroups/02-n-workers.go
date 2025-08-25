/*
The below code demonstrates the use of WaitGroups to synchronize multiple goroutines (workers).
Each worker prints a message when it finishes, and the main function waits for all to 
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
	numWorkers := 5

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d finished\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers completed")
}

// Output:

/*
➜  goConcurrency git:(main) ✗ go run 02-waitgroups/02-n-workers.go
Worker 1 finished
Worker 5 finished
Worker 3 finished
Worker 2 finished
Worker 4 finished
All workers completed
*/

