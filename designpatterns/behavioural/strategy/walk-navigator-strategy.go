package strategy

import "fmt"

type WalkNavigatorStrategy struct{}

func (c *WalkNavigatorStrategy) navigator(from float64, to float64) {
	fmt.Println("Walk Navigator Strategy navigating")
}
