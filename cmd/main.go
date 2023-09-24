package main

import (
	// Importe o pacote com as informações de configuração
	"crud-go/config"
	"crud-go/internal/app/repository"
	"crud-go/internal/app/service"
	"crud-go/internal/pkg/router"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Chame a função GetDBConfig para obter as informações de configuração do banco de dados
	dbUser, dbPassword, dbHost, dbPort, dbName, error := config.GetDBConfig()

	if error != nil {
		panic(error)
	}

	// Use as informações do banco de dados para configurar a conexão com o MySQL
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Abra a conexão com o MySQL
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Teste a conexão
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	log.Printf("Conexão com o MySQL estabelecida com sucesso!") // fixed line

	// Configuração do repositório (por exemplo, in-memory ou banco de dados)
	userRepository := repository.NewInMemoryUserRepository()

	// Configuração do serviço
	userService := service.NewUserService(userRepository)

	// Configuração do roteador
	r := router.SetupRouter(userService)

	// Inicia o servidor
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
