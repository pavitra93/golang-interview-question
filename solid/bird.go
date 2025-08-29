package solid

type bird struct {
	name     string
	color    string
	size     string
	birdType BirdType
}

func (b *bird) GetBirdType() BirdType {
	return b.birdType
}

func (b *bird) GetColor() string {
	return b.color
}

func (b *bird) GetSize() string {
	return b.size
}

func (b *bird) SetName(name string) {
	b.name = name
}

func (b *bird) SetColor(color string) {
	b.color = color
}

func (b *bird) SetSize(size string) {
	b.size = size
}

func (b *bird) SetBirdType(birdType BirdType) {
	b.birdType = birdType
}
