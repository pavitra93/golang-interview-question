package service

import (
	"fmt"

	"github.com/pavitra93/trello/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (service *UserService) AddUser(name string, email string) (int, error) {
	User, err := service.UserRepository.Create(name, email)
	if err != nil {
		fmt.Printf("Error creating user: %s\n", err)
		return 0, err
	}

	fmt.Printf("User created: %d\n", User.ID)
	return User.ID, err
}

func (service *UserService) ListUser() {
	Users, err := service.UserRepository.List()
	if err != nil {
		fmt.Printf("Error listing users: %s\n", err)
	}

	for _, User := range Users {
		fmt.Printf("User: %s\n", User.Name)
	}
}

func (service *UserService) SingleUser(id int) {
	User, err := service.UserRepository.GetUserByID(id)
	if err != nil {
		fmt.Printf("Error getting user: %s\n", err)
	}

	fmt.Printf("Single User: %s\n", User.Name)

}
