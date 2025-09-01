package models

type List struct {
	ID    int
	Name  string
	Cards []int
}

func NewList(id int, name string, Cards []int) *List {
	return &List{
		ID:    id,
		Name:  name,
		Cards: Cards,
	}
}
