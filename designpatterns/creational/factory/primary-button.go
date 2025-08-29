package factory

import "fmt"

type PrimaryButton struct {
	Type ButtonType
}

func (p *PrimaryButton) ClickButton() {
	fmt.Printf("Primary button type is %s\n", p.Type)
}
