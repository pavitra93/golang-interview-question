package repository

import (
	"errors"
	"math/rand"

	"github.com/pavitra93/trello/models"
)

type UserRepository struct {
	Users map[int]*models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Users: make(map[int]*models.User),
	}
}

func (UserRepository *UserRepository) Create(name string, email string) (*models.User, error) {
	User := &models.User{
		ID:    rand.Int(),
		Name:  name,
		Email: email,
	}

	UserRepository.Users[User.ID] = User
	return User, nil
}

func (UserRepository *UserRepository) List() ([]*models.User, error) {
	UserList := make([]*models.User, 0, len(UserRepository.Users))
	for _, user := range UserRepository.Users {
		UserList = append(UserList, user)
	}

	return UserList, nil
}

func (UserRepository *UserRepository) GetUserByID(id int) (*models.User, error) {
	User, ok := UserRepository.Users[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return User, nil
}
