package models

type Card struct {
	ID             int
	Name           string
	Desc           string
	AssignedUserID int
}

func NewCard(id int, name string, desc string, assignedUserID int) *Card {
	return &Card{
		ID:             id,
		Name:           name,
		Desc:           desc,
		AssignedUserID: assignedUserID,
	}
}
