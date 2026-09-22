package main

import "fmt"

type Instructor struct {
	FirstName string
	LastName  string
}

func main() {
	var lasse Instructor
	lasse = Instructor{FirstName: "Lasse", LastName: "Jensen"}

	paolo := Instructor{"Paolo", "Tell"}

	fmt.Println(lasse)
	fmt.Println(paolo)
}
