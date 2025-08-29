package prototype

type Rectangle struct {
	Red   float64
	Green float64
	Blue  float64
}

func (r *Rectangle) Clone() Shape {

	return &Rectangle{Red: r.Red, Green: r.Green, Blue: r.Blue}
}
