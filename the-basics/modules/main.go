package main

import (
	"fmt"
)

var packageVariable = "I am a package variable\n"

func init() {
	fmt.Println("A")
}

func init() {
	fmt.Println("B")
}

func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

func calculateTaxWithNames(price float32) (stateTax float32, cityTax float32) {
	stateTax = price * 0.09
	cityTax = price * 0.02
	return
}

func main() {
	stateTax, cityTax := calculateTax(100)
	stateTax2, _ := calculateTaxWithNames(100)

	fmt.Println(stateTax, cityTax)
	fmt.Println(stateTax2)

	printSomeData()
}
