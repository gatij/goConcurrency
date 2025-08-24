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