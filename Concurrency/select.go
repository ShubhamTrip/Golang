package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string) // Buffered channels
	ch2 := make(chan string)

	go one(ch1)
	go two(ch2)

	// No need for time.Sleep or timeout here
	select {
	case val1 := <-ch1:
		fmt.Println("Received from ch1:", val1)
		break
	case val2 := <-ch2:
		fmt.Println("Received from ch2:", val2)
		break
	default:
		fmt.Println("No data received from either channel")
		break
	}
}

func one(ch1 chan string) {
	ch1 <- "World"
}

func two(ch2 chan string) {
	ch2 <- "Hello"
}
