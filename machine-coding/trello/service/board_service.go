package service

import (
	"fmt"

	"github.com/pavitra93/trello/enums/privacy"
	"github.com/pavitra93/trello/models"
	"github.com/pavitra93/trello/repository"
)

type BoardService struct {
	BoardRepository *repository.BoardRepository
}

func (service *BoardService) CreateBoard(name string, desc string, privacy privacy.Privacy) (int, error) {
	Board, err := service.BoardRepository.CreateBoard(name, desc, privacy)
	if err != nil {
		fmt.Printf("CreateBoard error: %v\n", err)
		return 0, err
	}

	fmt.Printf("Created Board %v\n", Board)
	return Board.ID, nil
}

func (service *BoardService) SingleBoard(id int) {
	Board, err := service.BoardRepository.GetBoardByID(id)
	if err != nil {
		fmt.Printf("GetBoard error: %v\n", err)
	}

	fmt.Printf("Single Board %v\n", Board)
}

func (service *BoardService) DeleteBoard(id int) {
	Board, err := service.BoardRepository.GetBoardByID(id)
	if err != nil {
		fmt.Printf("GetBoard error: %v\n", err)
	}
	service.BoardRepository.DeleteBoardByID(Board.ID)
	fmt.Printf("Deleted Board %v\n", Board)
}

func (service *BoardService) ListBoards() ([]*models.Board, error) {
	boards, err := service.BoardRepository.GetAllBoard()
	if err != nil {
		fmt.Printf("ListBoards error: %v\n", err)
	}

	return boards, err
}

func (service *BoardService) AddListsByBoardID(id int, lists []int) {
	Board, err := service.BoardRepository.GetBoardByID(id)
	if err != nil {
		fmt.Printf("GetBoard error: %v\n", err)
	}

	Board.Lists = append(Board.Lists, lists...)
	fmt.Printf("Lists Added %v\n", Board.Lists)
}

func (service *BoardService) AddUsersByBoardID(id int, users []int) {
	Board, err := service.BoardRepository.GetBoardByID(id)
	if err != nil {
		fmt.Printf("Get Board error: %v\n", err)
	}

	Board.Users = append(Board.Users, users...)
	fmt.Printf("Users Added %v\n", Board.Users)
}
