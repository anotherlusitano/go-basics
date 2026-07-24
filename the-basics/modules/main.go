package main

import "epic.example/data"

var packageVariable = "I am a package variable\n"

func main() {
	const pi = 3.14

	var hello string = "Hello, World!\n"

	price := 49.9

	print(hello, price, data.Pi)

	printSomeData()
}
