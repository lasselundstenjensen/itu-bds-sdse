package coursetypes

import (
	"time"
)

type Instructor struct {
	FirstName string
	LastName  string
}

type Course struct {
	Name       string
	Instructor Instructor
	Date       time.Time
}
