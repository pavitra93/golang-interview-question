package solid

import "fmt"

type Parrot struct {
	bird
}

func (p *Parrot) Fly() {
	fmt.Printf("%s is Flying\n", p.birdType)
}
