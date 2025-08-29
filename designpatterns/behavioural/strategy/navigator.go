package strategy

type Navigator struct {
	navigatorStrategy NavigatorStrategy
}

func (n *Navigator) SetStrategy(strategy NavigatorStrategy) {
	n.navigatorStrategy = strategy
}

func (n *Navigator) Navigate(from float64, to float64) {
	n.navigatorStrategy.navigator(from, to)
}
