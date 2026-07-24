package data

import "fmt"

var Countries [10]string

var Slices []int

var Codes map[int][]string

func init() {
	Countries[0] = "Portugal"
	Countries[1] = "Brazil"
	Countries[9] = "USA"

	qty := len(Countries)

	fmt.Println("Countries saved", qty)
}
