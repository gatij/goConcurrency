# Go Concurrency Playground

A collection of simple, practical examples demonstrating concurrency in Go using goroutines. Each example is self-contained and designed to illustrate a specific concept or problem related to concurrent programming.

## Structure

- `01-goroutines/`
  - `01-simple-go-routine.go`: Launches a single goroutine.
  - `02-multiple-go-routine.go`: Runs multiple goroutines in parallel.
  - `03-scaled-go-routine.go`: Demonstrates scaling up goroutines (10 workers).
  - `problems/`: Contains concurrency exercises and their solutions:
    - `01-problem.go`: Launch 10 goroutines, each printing its own number.
    - `02-problem.go`: Launch 5 goroutines with different names.
    - `03-problem.go`: Goroutines print after random delays; output order varies.
    - `04-problem.go`: Compute squares in parallel using goroutines.
    - `05-problem.go`: Shows race between main and goroutine; output may be unpredictable.
    - `bonus-problem.go`: Simulates baking a cake with concurrent and sequential tasks.

## Getting Started

### Prerequisites
- [Go](https://golang.org/dl/) 1.23 or newer

### Running Examples
Navigate to the project directory and run any example:

```sh
cd 01-goroutines
# Run a simple goroutine example
go run 01-simple-go-routine.go

# Run a problem example
go run problems/01-problem.go
```

## What You'll Learn
- How to launch goroutines
- Synchronization basics (using `time.Sleep` for simplicity)
- How concurrency affects output order
- Practical problems and solutions in Go concurrency

## Author
Created by Gatij Jain

## License
MIT License
