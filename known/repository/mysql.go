package repository

import (
	"database/sql"
	"fmt"
	"go-cli-memo/known"
	"go-cli-memo/models"
)

type mysqlKnownRepository struct {
	Conn *sql.DB
}

func NewMysqlKnownRepository(Conn *sql.DB) known.Repository {

	return &mysqlKnownRepository{Conn}

}

func (mr *mysqlKnownRepository) Get() ([]*models.Known, error) {
	rows, err := mr.Conn.Query("SELECT * FROM known")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	result := make([]*models.Known, 0)
	for rows.Next() {
		var k = models.Known{}
		err = rows.Scan(&k.ID, &k.Word, &k.Description)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, &k)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return result, err
}

func (mr *mysqlKnownRepository) Store(k *models.Known) error {
	query := fmt.Sprintf("INSERT INTO known (Word, Description) VALUES('%s', '%s')", k.Word, k.Description)
	_, err := mr.Conn.Exec(query)
	if err != nil {
		panic(err.Error())
	}

	return err
}

func (mr *mysqlKnownRepository) Update(k *models.Known) error {
	query := fmt.Sprintf("UPDATE known SET Description = '%s' where Word = '%s'", k.Description, k.Word)
	_, err := mr.Conn.Exec(query)
	if err != nil {
		panic(err.Error())
	}

	return err
}

func (mr *mysqlKnownRepository) Delete(k *models.Known) error {
	query := fmt.Sprintf("DELETE FROM known WHERE Word = '%s'", k.Word)
	_, err := mr.Conn.Exec(query)
	if err != nil {
		panic(err.Error())
	}

	return err
}
