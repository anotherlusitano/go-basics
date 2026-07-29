package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func (i Instructor) Print() string {
	return fmt.Sprintf("%v, %v, (%d)\n", i.LastName, i.FirstName, i.Score)
}
