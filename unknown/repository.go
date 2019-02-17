package unknown

import "go-cli-memo/models"

type Repository interface {
	Get() ([]*models.Unknown, error)
	Store(u *models.Unknown) error
	Update(u1, u2 *models.Unknown) error
	Delete(u *models.Unknown) (int64, error)
}
