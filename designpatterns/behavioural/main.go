package main

import (
	"github.com/pavitra93/golang-interview-questions.git/designpatterns/behavioural/strategy"
)

func main() {
	navigator := &strategy.Navigator{}
	navigator.SetStrategy(&strategy.CarNavigatorStrategy{})
	navigator.Navigate(1.0, 2.0)

	navigator.SetStrategy(&strategy.WalkNavigatorStrategy{})
	navigator.Navigate(1.0, 2.0)
}
