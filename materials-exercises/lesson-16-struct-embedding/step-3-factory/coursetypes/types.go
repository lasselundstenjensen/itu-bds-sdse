package coursetypes

import (
	"fmt"
	"time"
)

type Instructor struct {
	FirstName string
	LastName  string
}

func (i Instructor) Hello() {
	fmt.Printf("Hello, I'm %s %s\n", i.FirstName, i.LastName)
}

func NewInstructor(firstname string, lastname string) Instructor {
	return Instructor{FirstName: firstname, LastName: lastname}
}

type Course struct {
	Name       string
	Instructor Instructor
	Date       time.Time
}

func (c Course) String() string {
	return fmt.Sprintf("Course name: %s", c.Name)
}

type Workshop struct {
	Course
	SignupDate time.Time
}

func NewWorkshop(courseName string, instructor Instructor, signupDate time.Time) Workshop {
	return Workshop{
		Course:     Course{Name: courseName, Instructor: instructor},
		SignupDate: signupDate,
	}
}
