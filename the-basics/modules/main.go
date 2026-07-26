package main

import (
	"fmt"
)

var packageVariable = "I am a package variable\n"

func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

func calculateTaxWithNames(price float32) (stateTax float32, cityTax float32) {
	stateTax = price * 0.09
	cityTax = price * 0.02
	return
}

func birthday(age int) {
	age = age + 1
}

func birthdayReal(age *int) {
	defer fmt.Println("Happy birthday!")

	if *age > 120 {
		// panic is just to exit the program, but it is not a good practice to use panic for error handling
		panic("You shall not live that long!")
	}

	fmt.Printf("Pointer: %v \nValue: %v \n", age, *age)

	*age++
}

func main() {
	// stateTax, _ := calculateTaxWithNames(100)
	// fmt.Println(stateTax)

	// will execute at the end of the function, even if there is a panic
	defer fmt.Println("Byeee")
	defer fmt.Println("The first to enter is the last to leave, like a stack!")

	age := 30
	birthday(age)
	fmt.Println(age)

	birthdayReal(&age)
	fmt.Println(age)

	printSomeData()
}
