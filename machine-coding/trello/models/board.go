package models

import (
	"github.com/pavitra93/trello/enums/privacy"
)

type Board struct {
	ID      int
	Name    string
	Desc    string
	Privacy privacy.Privacy
	Lists   []int
	Users   []int
}

func NewBoard(id int, name string, desc string, privacy privacy.Privacy, lists []int, users []int) *Board {
	return &Board{
		ID:      id,
		Name:    name,
		Desc:    desc,
		Privacy: privacy,
		Lists:   lists,
		Users:   users,
	}
}
