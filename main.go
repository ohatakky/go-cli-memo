package main

import (
	"database/sql"
	"fmt"
	_unknownCliDeliver "go-cli-memo/unknown/delivery/cli"
	_unknownRepo "go-cli-memo/unknown/repository"
	_unknownUsecase "go-cli-memo/unknown/usecase"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func init() {
	env_err := godotenv.Load()
	if env_err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dbUser := os.Getenv("dbUser")
	dbPass := os.Getenv("dbPass")
	dbName := os.Getenv("dbName")
	connection := fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName)

	db, err := sql.Open(`mysql`, connection)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	unknownRepo := _unknownRepo.NewMysqlUnknownRepository(db)
	unknownUsecase := _unknownUsecase.NewUnknownUsecase(unknownRepo)

	app := cli.NewApp()
	_unknownCliDeliver.NewUnknownCliHandler(app, unknownUsecase)
	app.Run(os.Args)
}
