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

	rustWS := data.NewWorkshop("Rust", max)

	fmt.Printf("%v\n", rustWS)

	fmt.Println("-------- Interfaces ---------")

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = rustWS.Course

	for _, course := range courses {
		fmt.Println(course)
	}

	var things []interface{}

	things = append(things, "34")
	things = append(things, 34)
	things = append(things, rustWS)

	things[2].(data.Workshop).SignUp()

	fmt.Println(things)
}
