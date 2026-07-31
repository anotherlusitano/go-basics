package data

import "fmt"

type Duration float32 // in hours

type Course struct {
	Id         int
	Name       string
	Slug       string
	Legacy     bool
	Duration   Duration
	Instructor Instructor
}

func New(instructor Instructor) Course {
	return Course{Id: 1, Name: "Go Fundamentals", Slug: "go-fundamentals", Legacy: false, Duration: 2.5, Instructor: instructor}
}

func (c Course) Print() string {
	return fmt.Sprintf("%v, %v, (%d)\n", c.Name, c.Slug, c.Id)
}
