package main

import (
	"fmt"
	"time"

	"itu.dk/course/coursetypes"
)

func main() {
	lasse := coursetypes.Instructor{
		FirstName: "Lasse",
		LastName:  "Jensen"}

	paolo := coursetypes.NewInstructor("Paolo", "Tell")
	fmt.Println(paolo)

	course := coursetypes.Course{}
	course.Name = "Go course"
	course.Instructor = lasse
	course.Date = time.Now()

	fmt.Println(lasse)
	fmt.Println(course)

	workshop := coursetypes.Workshop{
		Course:     coursetypes.Course{Name: "Workshop on Go", Date: time.Now()},
		SignupDate: time.Now(),
		Name:       "Workshop",
	}

	fmt.Println(workshop)
	fmt.Println(workshop.Name)
	fmt.Println(workshop.Course.Name)

	workshop.Course.Instructor = paolo
	workshop.Course.Instructor.Hello()

	fmt.Println()

	lasse.Hello()
	course.Instructor.Hello()

	lasse.FirstName = "Paolo"
	course.Instructor.Hello()
}
