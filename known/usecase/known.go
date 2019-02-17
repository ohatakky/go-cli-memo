package usecase

import (
	"go-cli-memo/known"
	"go-cli-memo/models"
	"go-cli-memo/unknown"
	"log"
)

type knownUsecase struct {
	knownRepo   known.Repository
	unknownRepo unknown.Repository
}

func NewKnownUsecase(kRepo known.Repository, ukRepo unknown.Repository) known.Usecase {
	return &knownUsecase{
		knownRepo:   kRepo,
		unknownRepo: ukRepo,
	}
}

func (ku *knownUsecase) Get() ([]*models.Known, error) {

	res, err := ku.knownRepo.Get()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ku *knownUsecase) Store(k *models.Known) error {
	// TODO : トランザクション張る。
	var u = models.Unknown{Word: k.Word}
	n, err2 := ku.unknownRepo.Delete(&u)

	if err2 != nil {
		return err2
	}
	if n == 0 { // delete件数が0ならreturn
		log.Fatal("該当するunknownなし")
	}

	err := ku.knownRepo.Store(k)
	if err != nil {
		return err
	}

	return nil
}

func (ku *knownUsecase) Update(k *models.Known) error {
	err := ku.knownRepo.Update(k)
	if err != nil {
		return err
	}

	return nil
}

func (ku *knownUsecase) Delete(k *models.Known) error {
	err := ku.knownRepo.Delete(k)
	if err != nil {
		return err
	}

	return nil
}
