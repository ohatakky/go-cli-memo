package cli

import (
	"fmt"
	"go-cli-memo/known"
	"go-cli-memo/models"
	"os"

	"github.com/urfave/cli"
)

type CliKnownHandler struct {
	kUse known.Usecase
}

func NewKnownCliHandler(ku known.Usecase) {
	handler := &CliKnownHandler{
		kUse: ku,
	}
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "known",
			Usage: "known CRUD",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "update, u",
					Usage: "update known word",
				},
			},
			Action: func(c *cli.Context) error {
				switch len(c.Args()) {
				case 0:
					k, err := handler.get()
					for _, v := range k {
						fmt.Println(v.Word, v.Description)
					}
					if err != nil {
						return err
					}
				case 2:
					var k = models.Known{Word: c.Args()[0], Description: c.Args()[1]}
					err := handler.store(&k)
					if err != nil {
						return err
					}
				case 3:
					if c.Bool("u") {
						var k = models.Known{Word: c.Args()[0], Description: c.Args()[2]}
						err := handler.update(&k)
						if err != nil {
							return err
						}
					}
				}
				return nil
			},
		},
	}

	app.CommandNotFound = func(ctx *cli.Context, command string) {}

	app.Run(os.Args)
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
