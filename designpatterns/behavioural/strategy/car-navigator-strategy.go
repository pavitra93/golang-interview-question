package strategy

import "fmt"

type CarNavigatorStrategy struct{}

func (c *CarNavigatorStrategy) navigator(from float64, to float64) {
	fmt.Println("Car Navigator Strategy navigating")
}
