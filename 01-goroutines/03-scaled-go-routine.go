package main

import (
	"fmt"
	"time"
)

func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    for i := 0; i < 10; i++ {
        go worker(i)
    }
    time.Sleep(2 * time.Second)
}

// You can start thousands of goroutines—much cheaper than OS threads!

// This example starts 10 goroutines, each simulating a worker doing some work.
// The main function sleeps for 2 seconds to allow all workers to complete before exiting.

/*
➜  goConcurrency go run 01-goroutines/03-scaled-go-routine.go  
Worker 0 starting
Worker 1 starting
Worker 2 starting
Worker 3 starting
Worker 4 starting
Worker 5 starting
Worker 6 starting
Worker 7 starting
Worker 8 starting
Worker 9 starting
Worker 0 done
Worker 1 done
Worker 2 done
Worker 3 done
Worker 4 done
Worker 5 done
Worker 6 done
Worker 7 done
Worker 8 done
Worker 9 done
*/