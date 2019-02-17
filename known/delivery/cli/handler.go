package cli

import (
	"go-cli-memo/known"
	"go-cli-memo/models"

	"github.com/urfave/cli"
)

type CliKnownHandler struct {
	kUse known.Usecase
}

func NewKnownCliHandler(app *cli.App, ku known.Usecase) {
	handler := &CliKnownHandler{
		kUse: ku,
	}
	// TODO : コマンドとフラグでswitchさせる
	handler.get()
}

func (h *CliKnownHandler) get() ([]*models.Known, error) {

	known, err := h.kUse.Get()
	if err != nil {
		return nil, err
	}

	return known, nil
}

func (h *CliKnownHandler) store(u *models.Known) error {
	err := h.kUse.Store(u)
	if err != nil {
		return err
	}

	return nil
}

func (h *CliKnownHandler) update(k *models.Known) error {
	err := h.kUse.Update(k)
	if err != nil {
		return err
	}

	return nil
}

func (h *CliKnownHandler) delete(k *models.Known) error {
	err := h.kUse.Delete(k)
	if err != nil {
		return err
	}

	return nil
}
