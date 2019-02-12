package usecase

import (
	"go-cli-memo/models"
	"go-cli-memo/unknown"
)

type unknownUsecase struct {
	unknownRepo unknown.Repository
}

func NewUnknownUsecase(ukRepo unknown.Repository) unknown.Usecase {
	return &unknownUsecase{unknownRepo: ukRepo}
}

func (uu *unknownUsecase) Get() ([]*models.Unknown, error) {

	res, err := uu.unknownRepo.Get()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uu *unknownUsecase) Store(u *models.Unknown) error {

	err := uu.unknownRepo.Store(u)
	if err != nil {
		return err
	}

	return nil
}
