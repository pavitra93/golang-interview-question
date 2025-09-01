package repository

import (
	"errors"
	"math/rand"

	"github.com/pavitra93/trello/enums/privacy"
	"github.com/pavitra93/trello/models"
)

type BoardRepository struct {
	Boards map[int]*models.Board
}

func NewBoardRepository() *BoardRepository {
	return &BoardRepository{
		Boards: make(map[int]*models.Board),
	}
}

func (repository *BoardRepository) CreateBoard(name string, desc string, privacy privacy.Privacy) (*models.Board, error) {
	Board := &models.Board{
		ID:      rand.Int(),
		Name:    name,
		Desc:    desc,
		Privacy: privacy,
	}

	repository.Boards[Board.ID] = Board

	return Board, nil
}

func (repository *BoardRepository) GetBoardByID(id int) (*models.Board, error) {
	Board, ok := repository.Boards[id]
	if !ok {
		return nil, errors.New("board not found")
	}

	return Board, nil
}

func (repository *BoardRepository) GetAllBoard() ([]*models.Board, error) {
	Boards := make([]*models.Board, 0)
	for _, board := range repository.Boards {
		Boards = append(Boards, board)
	}

	return Boards, nil
}

func (repository *BoardRepository) DeleteBoardByID(id int) {
	delete(repository.Boards, id)
}
