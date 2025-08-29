package solid

import "fmt"

type Peguin struct {
	NonCommonAttribute string
	bird
}

func (p *Peguin) SetNonCommonAttribute(attr string) {
	p.NonCommonAttribute = attr
}

func (p *Peguin) GetNonCommonAttribute() string {
	return p.NonCommonAttribute
}

func (p *Peguin) Swim() {
	fmt.Printf("%s is Swimming\n", p.birdType)

}
