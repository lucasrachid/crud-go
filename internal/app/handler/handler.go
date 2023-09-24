package handler

import (
	"crud-go/internal/app/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	// Implemente a lógica para criar um usuário.
}

func (h *UserHandler) GetUser(c *gin.Context) {
	// Implemente a lógica para obter um usuário.
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	// Implemente a lógica para atualizar um usuário.
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	// Implemente a lógica para excluir um usuário.
}
