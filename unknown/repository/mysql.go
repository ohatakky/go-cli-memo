package repository

import (
	"database/sql"
	"go-cli-memo/models"
)

type mysqlUnknownRepository struct {
	Conn *sql.DB
}

func Get() (*models.Unknown, error) {

	// 仮
	res := &models.Unknown{}
	return res, nil
}
