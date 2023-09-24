package router

import (
	"crud-go/internal/app/handler"
	"crud-go/internal/app/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userService *service.UserService) *gin.Engine {
	router := gin.Default()

	// Crie uma instância do handler de usuário.
	userHandler := handler.NewUserHandler(userService)

	// Use o middleware em um grupo de rotas específico.
	v1 := router.Group("/api/v1")
	{
		v1.Use(handler.CustomMiddleware()) // Adicione o middleware aqui.
		v1.POST("/users", userHandler.CreateUser)
		v1.GET("/users/:id", userHandler.GetUser)
		v1.PUT("/users/:id", userHandler.UpdateUser)
		v1.DELETE("/users/:id", userHandler.DeleteUser)
	}

	return router
}
