package main

import (
	"fmt"
	"time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}

func printLetters() {
    for ch := 'a'; ch <= 'c'; ch++ {
        fmt.Printf("%c\n", ch)
    }
}

func main() {
    go printNumbers()
    go printLetters()
    time.Sleep(1 * time.Second)
    fmt.Println("Done")
}

// Output order is unpredictable— a sign of concurrency!

/*
➜  goConcurrency go run 01-goroutines/02-multiple-go-routine.go
1
2
3
4
5
a
b
c
Done
➜  goConcurrency go run 01-goroutines/02-multiple-go-routine.go
a
b
c
1
2
3
4
5
Done
*/