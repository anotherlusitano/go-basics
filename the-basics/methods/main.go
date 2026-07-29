package main

import "methods.com/data"

func main() {
	max := data.Instructor{Id: 3, LastName: "Ronaldo"}
	max.FirstName = "Cristiano"

	print(max.Print())
}
