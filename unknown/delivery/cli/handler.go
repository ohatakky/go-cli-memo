package cli

import (
	"fmt"
	"go-cli-memo/models"
	"go-cli-memo/unknown"

	"github.com/urfave/cli"
)

type CliUnknownHandler struct {
	Uu unknown.Usecase
}

func NewUnknownCliHandler(app *cli.App, uu unknown.Usecase) {
	handler := &CliUnknownHandler{
		Uu: uu,
	}
	app.Name = "unknown"
	app.Usage = "get unknown list"
	app.Action = func(c *cli.Context) error {
		// TODO : 4コマンド実装。場合分けする。
		u, err := handler.get()
		for _, v := range u {
			fmt.Println(*v)
		}
		return err
	}
}

func (h *CliUnknownHandler) get() ([]*models.Unknown, error) {

	unknown, err := h.Uu.Get()
	if err != nil {
		return nil, err
	}

	return unknown, nil
}
