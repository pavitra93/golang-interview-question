package factory

type ButtonFactory struct{}

func (b *ButtonFactory) CreateButton(Type ButtonType) Button {
	switch Type {
	case "primary":
		return &PrimaryButton{Type: Type}
	case "secondary":
		return &SecondaryButton{Type: Type}
	default:
		return nil
	}
}
