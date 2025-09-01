package repository

import (
	"errors"
	"math/rand"

	"github.com/pavitra93/trello/models"
)

type CardRepository struct {
	Cards map[int]*models.Card
}

func NewCardRepository() *CardRepository {
	return &CardRepository{
		Cards: make(map[int]*models.Card),
	}
}

func (repository *CardRepository) CreateCard(name string, desc string) (*models.Card, error) {
	Card := &models.Card{
		ID:   rand.Int(),
		Name: name,
		Desc: desc,
	}

	repository.Cards[Card.ID] = Card

	return Card, nil
}

func (repository *CardRepository) GetCardByID(id int) (*models.Card, error) {
	Card, ok := repository.Cards[id]
	if !ok {
		return nil, errors.New("card not found")
	}

	return Card, nil
}

func (repository *CardRepository) GetAllCard() ([]*models.Card, error) {
	Cards := make([]*models.Card, 0)
	for _, card := range repository.Cards {
		Cards = append(Cards, card)
	}

	return Cards, nil
}

func (repository *CardRepository) DeleteCardByID(id int) {
	delete(repository.Cards, id)
}
