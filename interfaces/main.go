package main

import (
	"fmt"
	shapes "github.com/pavitra93/golang-interview-questions.git/interfaces/shapes"
)

func main() {
	fmt.Println("Hello Interfaces")

	c := &shapes.Circle{Radius: 5.0}
	fmt.Println(c.Area())
	fmt.Println(c.Circumference())

	r := &shapes.Rectangle{Length: 5.0, Width: 10.0}
	fmt.Println(r.Area())
	fmt.Println(r.Circumference())

}
