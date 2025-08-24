package main

import "fmt"
import "time"

func sayHello() {
    fmt.Println("Hello from goroutine!")
}

func main() {
    go sayHello() // Starts a new goroutine
    time.Sleep(1 * time.Second) // Wait for goroutine to finish
    fmt.Println("Main exited")
}

// Notice main() might exit before sayHello() executes! Therefore used time.Sleep() to wait.

/*
➜  goConcurrency go run 01-goroutines/01-simple-go-routine.go  
Hello from goroutine!
Main exited
*/