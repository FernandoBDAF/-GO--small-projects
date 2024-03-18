package main

import (
	"fmt"
	"structs/model"
	"time"
)

func main() {
	max := model.Instructor{Id: 1, FirstName: "Max", LastName: "Mustermann", Score: 100}

	kyle := model.NewInstructor(2, "Kyle", "Kowalski", 90)

	goCourse := model.Course{Id: 1, Name: "Go for Beginners", Slug: "go-for-beginners", Legacy: false,
		Duration: 4.5, Instructor: kyle}

	swiftWS := model.NewWorkshop("Swift Workshop", max, time.Now())

	println(max.Print())
	println(goCourse.String())
	fmt.Println(swiftWS)

	var courses [2]model.Signable
	courses[0] = goCourse
	courses[1] = swiftWS

	for _, c := range courses {
		fmt.Println(c)
	}
}
