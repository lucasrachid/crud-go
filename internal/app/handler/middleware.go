package handler

import (
	"github.com/gin-gonic/gin"
)

func CustomMiddleware() gin.HandlerFunc {
	// Este é um exemplo de middleware que pode ser personalizado de acordo com as necessidades do seu aplicativo.
	return func(c *gin.Context) {
		// Código a ser executado antes do tratamento da requisição.

		// Chame c.Next() para passar a requisição para o próximo middleware ou tratamento.
		c.Next()

		// Código a ser executado depois do tratamento da requisição.
	}
}
