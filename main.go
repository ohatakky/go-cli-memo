package main

import (
	"database/sql"
	"fmt"
	_knownCliDeliver "go-cli-memo/known/delivery/cli"
	_knownRepo "go-cli-memo/known/repository"
	_knownUsecase "go-cli-memo/known/usecase"
	_unknownCliDeliver "go-cli-memo/unknown/delivery/cli"
	_unknownRepo "go-cli-memo/unknown/repository"
	_unknownUsecase "go-cli-memo/unknown/usecase"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func init() {
	abs_path, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}
	env_err := godotenv.Load(fmt.Sprintf("%s/.env", abs_path))
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
	knownRepo := _knownRepo.NewMysqlKnownRepository(db)

	unknownUsecase := _unknownUsecase.NewUnknownUsecase(unknownRepo)
	knownUsecase := _knownUsecase.NewKnownUsecase(knownRepo, unknownRepo)

	app := cli.NewApp() // CLIパッケージ
	_unknownCliDeliver.NewUnknownCliHandler(app, unknownUsecase)
	_knownCliDeliver.NewKnownCliHandler(app, knownUsecase)
	app.Run(os.Args)
}
