package main

import (
	"fmt"

	"methods.com/data"
)

func main() {
	max := data.Instructor{Id: 3, LastName: "Ronaldo"}
	max.FirstName = "Cristiano"

	homer := data.NewIntructor("Homer", "Simpson")

	print(max.Print())
	print(homer.Print())

	goCourse := data.New(homer)

	print(goCourse.Print())

	fmt.Println(goCourse)
}
