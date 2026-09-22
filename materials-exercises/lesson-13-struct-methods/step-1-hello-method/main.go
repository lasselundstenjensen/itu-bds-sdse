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

	course := coursetypes.Course{}
	course.Name = "Go course"
	course.Instructor = lasse
	course.Date = time.Now()

	fmt.Println(lasse)
	fmt.Println(course)

	fmt.Println()

	lasse.Hello()
	course.Instructor.Hello()
}
