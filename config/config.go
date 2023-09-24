package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvError struct {
	Msg string
}

func (e *EnvError) Error() string {
	return e.Msg
}

// GetDBConfig obtém as informações do banco de dados do arquivo de configuração
func GetDBConfig() (string, string, string, string, string, error) {
	// Carregue as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Erro ao carregar arquivo .env:", err)
		return "", "", "", "", "", err
	}

	// Obtenha as variáveis de ambiente
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		error := &EnvError{
			Msg: "Dados do Environtment estão vazios.",
		}
		return "", "", "", "", "", error
	}

	return dbUser, dbPassword, dbHost, dbPort, dbName, nil
}
