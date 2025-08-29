package solid

import "fmt"

type Duck struct {
	bird
}

func (d *Duck) Fly() {
	fmt.Printf("%s can Fly\n", d.birdType)
}
func (d *Duck) Swim() {
	fmt.Printf("%s can Swim\n", d.birdType)
}
