package data

import "time"

type Workshop struct {
	// Id int
	Course
	Date time.Time
}

func (c Workshop) SignUp() bool {
	return true
}

func NewWorkshop(name string, instructor Instructor) Workshop {
	w := Workshop{}
	w.Name = name
	// w.Id = 99
	// w.Instructor.Id = 99
	w.Instructor = instructor
	return w
}
