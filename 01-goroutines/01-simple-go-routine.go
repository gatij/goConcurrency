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