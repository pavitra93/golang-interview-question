package repository

import (
	"errors"
	"math/rand"

	"github.com/pavitra93/trello/models"
)

type ListRepository struct {
	Lists map[int]*models.List
}

func NewListRepository() *ListRepository {
	return &ListRepository{
		Lists: make(map[int]*models.List),
	}
}

func (repository *ListRepository) CreateList(name string) (*models.List, error) {
	List := &models.List{
		ID:   rand.Int(),
		Name: name,
	}

	repository.Lists[List.ID] = List

	return List, nil
}

func (repository *ListRepository) GetListByID(id int) (*models.List, error) {
	List, ok := repository.Lists[id]
	if !ok {
		return nil, errors.New("list not found")
	}

	return List, nil
}

func (repository *ListRepository) GetAllList() ([]*models.List, error) {
	Lists := make([]*models.List, 0)
	for _, list := range repository.Lists {
		Lists = append(Lists, list)
	}

	return Lists, nil
}

func (repository *ListRepository) DeleteListByID(id int) {
	delete(repository.Lists, id)
}
