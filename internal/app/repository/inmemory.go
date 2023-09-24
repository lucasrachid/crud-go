package repository

import "crud-go/internal/app/model"

type InMemoryUserRepository struct {
	users map[int]*model.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[int]*model.User),
	}
}

func (r *InMemoryUserRepository) Create(user *model.User) error {
	// Implemente a criação aqui.
	return nil
}

func (r *InMemoryUserRepository) Read(id int) (*model.User, error) {
	// Implemente a leitura aqui.
	return nil, nil
}

func (r *InMemoryUserRepository) Update(id int, user *model.User) error {
	// Implemente a atualização aqui.
	return nil
}

func (r *InMemoryUserRepository) Delete(id int) error {
	// Implemente a exclusão aqui.
	return nil
}
