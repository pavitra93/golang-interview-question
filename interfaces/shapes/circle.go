package shapes

import "math"

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c *Circle) Circumference() float64 {
	return 2 * c.Radius * math.Pi
}
