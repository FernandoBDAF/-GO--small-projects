package model

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func NewInstructor(id int, firstName string, lastName string, score int) Instructor {
	return Instructor{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Score:     score,
	}
}

func (i Instructor) Print() string { 
	return fmt.Sprintf("Instructor: %s, %s (score: %v)", i.LastName, i.FirstName, i.Score)
}