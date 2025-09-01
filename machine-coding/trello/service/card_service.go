package service

import (
	"fmt"
	"slices"

	"github.com/pavitra93/trello/models"
	"github.com/pavitra93/trello/repository"
)

type CardService struct {
	CardRepository  *repository.CardRepository
	BoardRepository *repository.BoardRepository
	ListRepository  *repository.ListRepository
}

func (service *CardService) CreateCard(name string, desc string) (int, error) {
	Card, err := service.CardRepository.CreateCard(name, desc)
	if err != nil {
		fmt.Printf("Create Card error: %v\n", err)
		return 0, err
	}

	fmt.Printf("Created Card %v\n", Card)
	return Card.ID, nil
}

func (service *CardService) SingleCard(id int) {
	Card, err := service.CardRepository.GetCardByID(id)
	if err != nil {
		fmt.Printf("Get Card by ID error: %v\n", err)
	}

	fmt.Printf("Single Card %v\n", Card)
}

func (service *CardService) DeleteCard(id int) {
	Card, err := service.CardRepository.GetCardByID(id)
	if err != nil {
		fmt.Printf("Get Card error: %v\n", err)
	}
	service.CardRepository.DeleteCardByID(Card.ID)
	fmt.Printf("Deleted Card %v\n", Card)
}

func (service *CardService) ListCard() ([]*models.Card, error) {
	cards, err := service.CardRepository.GetAllCard()
	if err != nil {
		fmt.Printf("ListBoards error: %v\n", err)
	}

	return cards, err
}

func (service *CardService) AssignUserByCardID(id int, userId int) {
	Card, err := service.CardRepository.GetCardByID(id)
	if err != nil {
		fmt.Printf("Get CCard error: %v\n", err)
	}

	Card.AssignedUserID = userId
	fmt.Printf("Assigned User to Card %v\n", Card.AssignedUserID)
}

func (service *CardService) UnAssignUserByCardID(id int, userId int) {
	Card, err := service.CardRepository.GetCardByID(id)
	if err != nil {
		fmt.Printf("Get Card error: %v\n", err)
	}

	if Card.AssignedUserID != userId {
		fmt.Printf("User to card mismatch. Cannot unassign  %v\n", userId)
		return
	}

	Card.AssignedUserID = 0
	fmt.Printf("Unassigned User to Card %v\n", Card.AssignedUserID)
}

func (service *CardService) MoveCardAcrossListInSameBoard(boardId int, fromListId int, toListId int, cardId int) (bool, error) {
	Board, err := service.BoardRepository.GetBoardByID(boardId)
	if err != nil {
		fmt.Printf("Get Board error: %v\n", err)
		return false, err
	}

	if !slices.Contains(Board.Lists, fromListId) || !slices.Contains(Board.Lists, toListId) {
		fmt.Printf("list not found in board %v\n", boardId)
		return false, nil
	}

	fromList, err := service.ListRepository.GetListByID(fromListId)
	if err != nil {
		fmt.Printf("List not found error: %v\n", err)
		return false, nil
	}
	if !slices.Contains(fromList.Cards, cardId) {
		fmt.Printf("Card not found %v\n", cardId)
		return false, nil
	}

	toList, err := service.ListRepository.GetListByID(toListId)
	if err != nil {
		fmt.Printf("Destination List not found error: %v\n", err)
	}

	for idx, id := range fromList.Cards {
		if cardId == id {
			fromList.Cards = append(fromList.Cards[:idx], fromList.Cards[idx+1:]...)
			break
		}
	}

	toList.Cards = append(toList.Cards, cardId)
	return true, nil

}
