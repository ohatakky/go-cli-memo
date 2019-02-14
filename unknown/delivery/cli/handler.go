package cli

import (
	"fmt"
	"go-cli-memo/models"
	"go-cli-memo/unknown"
	"log"

	"github.com/urfave/cli"
)

type CliUnknownHandler struct {
	ukUse unknown.Usecase
}

func NewUnknownCliHandler(app *cli.App, uu unknown.Usecase) {
	handler := &CliUnknownHandler{
		ukUse: uu,
	}
	app.Action = func(c *cli.Context) error {
		if len(c.Args()) > 2 || len(c.Args()) == 0 {
			log.Fatal("arguments error")
		}

		if c.Args()[0] != "unknown" {
			log.Fatal("arguments error")
		}

		switch len(c.Args()) {
		case 1:
			u, err := handler.get()
			for _, v := range u {
				fmt.Println(v.Word)
			}
			if err != nil {
				return err
			}
		case 2:
			var u models.Unknown
			u.Word = c.Args()[1]
			err := handler.store(&u)
			return err
		}

		return nil
	}
}

func (h *CliUnknownHandler) get() ([]*models.Unknown, error) {

	unknown, err := h.ukUse.Get()
	if err != nil {
		return nil, err
	}

	return unknown, nil
}

func (h *CliUnknownHandler) store(u *models.Unknown) error {
	err := h.ukUse.Store(u)
	if err != nil {
		return err
	}

	return nil
}
