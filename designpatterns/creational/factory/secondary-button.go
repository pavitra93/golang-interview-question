package factory

import "fmt"

type SecondaryButton struct {
	Type ButtonType
}

func (p *SecondaryButton) ClickButton() {
	fmt.Printf("Secondary button type is %s\n", p.Type)
}
