/*
Title: "Race with the Main"
Task: Show what happens when you spin off a goroutine and let main exit immediately without synchronization. 
      Can you ensure output is predictable?
*/

package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func main() {
	go sayHello() // Starts a new goroutine
	fmt.Println("Main exited")
}

// OUTPUT: 
// Often "Main exited" prints before "Hello from goroutine!" due to main exiting first.
/*
➜  goConcurrency go run 01-goroutines/problems/05-problem.go
Main exited
Hello from goroutine!
➜  goConcurrency go run 01-goroutines/problems/05-problem.go
Main exited
➜  goConcurrency go run 01-goroutines/problems/05-problem.go
Main exited
➜  goConcurrency go run 01-goroutines/problems/05-problem.go
Main exited
Hello from goroutine!
➜  goConcurrency go run 01-goroutines/problems/05-problem.go
Main exited
*/
