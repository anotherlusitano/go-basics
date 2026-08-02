package main

import (
	"fmt"
	"time"
)

func printMessage(text string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}
}

// main goroutine
func main() {
	go printMessage("Go! Go!")
	go printMessage("Go! Java!")
	printMessage("Go! Rust!")

	// if there is nothing stoping the main goroutine,
	// it will exit before the other goroutines finish executing
	// go printMessage("Go! Rust!")
}
