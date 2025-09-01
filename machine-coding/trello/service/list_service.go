package service

import (
	"fmt"

	"github.com/pavitra93/trello/models"
	"github.com/pavitra93/trello/repository"
)

type ListService struct {
	ListRepository *repository.ListRepository
}

func (service *ListService) CreateList(name string) (int, error) {
	List, err := service.ListRepository.CreateList(name)
	if err != nil {
		fmt.Printf("Create List error: %v\n", err)
		return 0, err
	}

	fmt.Printf("Created List %v\n", List)
	return List.ID, nil
}

func (service *ListService) SingleList(id int) {
	List, err := service.ListRepository.GetListByID(id)
	if err != nil {
		fmt.Printf("Get List by ID error: %v\n", err)
	}

	fmt.Printf("Single List %v\n", List)
}

func (service *ListService) DeleteList(id int) {
	List, err := service.ListRepository.GetListByID(id)
	if err != nil {
		fmt.Printf("GetBoard error: %v\n", err)
	}
	service.ListRepository.DeleteListByID(List.ID)
	fmt.Printf("Deleted List %v\n", List)
}

func (service *ListService) ListLists() ([]*models.List, error) {
	lists, err := service.ListRepository.GetAllList()
	if err != nil {
		fmt.Printf("Listing Lists error: %v\n", err)
	}

	return lists, err
}

func (service *ListService) AddCardsByListID(id int, cards []int) {
	List, err := service.ListRepository.GetListByID(id)
	if err != nil {
		fmt.Printf("Get List error: %v\n", err)
	}

	List.Cards = append(List.Cards, cards...)
	fmt.Printf("Cards Added to List %v\n", List.Cards)
}
