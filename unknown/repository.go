package unknown

import "go-cli-memo/models"

type Repository interface {
	Get() ([]*models.Unknown, error)
	// Store()
}
