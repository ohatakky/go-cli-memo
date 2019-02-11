package main

import (
	"database/sql"
	"fmt"
	_unknownRepo "go-cli-memo/unknown/repository"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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
	res, err := unknownRepo.Get()
	for _, v := range res {
		fmt.Println(*v)
	}
	// unknownUse := _unknownUcase.NewUnknownUsecase(unknownRepo)
	// TODO : Delievery
}
