package repository

import "crud-go/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	Read(id int) (*model.User, error)
	Update(id int, user *model.User) error
	Delete(id int) error
}
