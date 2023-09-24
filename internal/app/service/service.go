package service

import (
	"crud-go/internal/app/model"
	"crud-go/internal/app/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		userRepository: repo,
	}
}

func (s *UserService) CreateUser(user *model.User) error {
	// Implemente a lógica de criação aqui.
	return nil
}

func (s *UserService) GetUser(id int) (*model.User, error) {
	// Implemente a lógica de leitura aqui.
	return nil, nil
}

func (s *UserService) UpdateUser(id int, user *model.User) error {
	// Implemente a lógica de atualização aqui.
	return nil
}

func (s *UserService) DeleteUser(id int) error {
	// Implemente a lógica de exclusão aqui.
	return nil
}
