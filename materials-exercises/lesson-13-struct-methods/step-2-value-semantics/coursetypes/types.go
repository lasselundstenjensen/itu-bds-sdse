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

type Course struct {
	Name       string
	Instructor Instructor
	Date       time.Time
}
