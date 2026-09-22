package main

import "fmt"

type distance float32

func main() {
	var miles distance
	miles = 20
	km := miles.ToKm()

	fmt.Println(km)
}

func (d distance) ToKm() distance {
	return d * 1.6
}
