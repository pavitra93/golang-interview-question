package solid

import "fmt"

type Eagle struct {
	bird // Inheritance
}

func (e *Eagle) Fly() {
	fmt.Printf("%s is Flying\n", e.birdType)
}
