package usecase

import (
	"go-cli-memo/models"
	"go-cli-memo/unknown"
)

type unknownUsecase struct {
	unknownRepo unknown.Repository
}

func NewUnknownUsecase(uk unknown.Repository) unknown.Usecase {
	return &unknownUsecase{unknownRepo: uk}
}

func (uk *unknownUsecase) Get() ([]*models.Unknown, error) {

	res, err := uk.unknownRepo.Get()
	if err != nil {
		return nil, err
	}

	return res, nil
}
