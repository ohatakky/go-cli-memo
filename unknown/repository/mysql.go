package repository

import (
	"database/sql"
	"fmt"
	"go-cli-memo/models"
	"go-cli-memo/unknown"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlUnknownRepository struct {
	Conn *sql.DB
}

func NewMysqlUnknownRepository(Conn *sql.DB) unknown.Repository {

	return &mysqlUnknownRepository{Conn}
}

func (mr *mysqlUnknownRepository) Get() ([]*models.Unknown, error) {
	rows, err := mr.Conn.Query("SELECT * FROM unknown")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	result := make([]*models.Unknown, 0)
	for rows.Next() {
		var uk = models.Unknown{}
		err = rows.Scan(&uk.ID, &uk.Word)
		if err != nil {
			panic(err.Error())
		}

		result = append(result, &uk)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return result, err
}

func (mr *mysqlUnknownRepository) Store(u *models.Unknown) error {
	insert_query := fmt.Sprintf("INSERT INTO unknown (Word) VALUES('%s')", u.Word)
	_, err := mr.Conn.Exec(insert_query)
	if err != nil {
		panic(err.Error())
	}

	return err
}
