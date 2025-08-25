/*
The below code demonstrates the use of WaitGroups to synchronize multiple goroutines performing computations.
Each goroutine computes the square of a number and signals completion using wg.Done().
The main function waits for all goroutines to finish using wg.Wait() before printing a final message.

This ensures that the main program does not exit before the goroutines finish their execution.
*/

/*
Compare this snippet from 01-goroutines/problems/04-problem.go:
The below code demonstrates launching multiple goroutines to perform computations in parallel.
Each goroutine computes the square of a number and prints the result along with the time taken.

The main function launches these goroutines and uses time.Sleep to wait for their completion.
This approach can lead to unpredictable behavior if the sleep duration is not sufficient.
*/

package main

import (
	"fmt"
	"sync"
)

func square(num int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Square of %d: %d\n", num, num*num)
}

func main() {
    nums := []int{2, 4, 6}
    var wg sync.WaitGroup
    for _, n := range nums {
        wg.Add(1)
        go square(n, &wg)
    }
    wg.Wait()
    fmt.Println("Squares computed")
}

// Output:
/*
➜  goConcurrency git:(main) ✗ go run 02-waitgroups/03-wg-and-computation.go
Square of 6: 36
Square of 2: 4
Square of 4: 16
Squares computed
*/