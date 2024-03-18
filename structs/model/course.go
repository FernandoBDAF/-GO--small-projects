package model

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

func (c Course) SignUp() bool {
	return true
}

func (c Course) String() string { 
	return fmt.Sprintf("--- %v --- (%v hrs)\n by %v", c.Name, c.Duration, c.Instructor.Print())
}