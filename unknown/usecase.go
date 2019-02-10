package unknown

import "go-cli-memo/models"

type Usecase interface {
	Get() (*models.Unknown, error)
	Store()
}
