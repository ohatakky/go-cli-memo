package known

import "go-cli-memo/models"

type Usecase interface {
	Get() ([]*models.Known, error)
	Store(k *models.Known) error
	Update(k *models.Known) error
	Delete(k *models.Known) error
}
