package model

import "time"

type Workshop struct {
	Course
	Instructor
	Date time.Time
}

func (c Workshop) SignUp() bool {
	return true
}

func NewWorkshop(name string, instructor Instructor, date time.Time) Workshop {
	w := Workshop{}
	w.Name = name
	w.Instructor = instructor
	w.Date = date
	return w
}