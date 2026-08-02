package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}

	channel <- "Done"
}

// main goroutine
func main() {
	channel := make(chan string)
	// can't use this without make:
	// var channel chan string

	go printMessage("Go! Go!", channel)

	response := <-channel

	fmt.Println(response)
}
