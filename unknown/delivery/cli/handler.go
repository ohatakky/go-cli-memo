package cli

import (
	"fmt"
	"go-cli-memo/models"
	"go-cli-memo/unknown"
	"log"
	"os"

	"github.com/urfave/cli"
)

type CliUnknownHandler struct {
	ukUse unknown.Usecase
}

func NewUnknownCliHandler(uu unknown.Usecase) {
	handler := &CliUnknownHandler{
		ukUse: uu,
	}
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "unknown",
			Usage: "unknown CRUD",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "update, u",
					Usage: "update unknown word",
				},
				cli.BoolFlag{
					Name:  "delete, d",
					Usage: "delete unknown word",
				},
			},
			Action: func(c *cli.Context) error {
				switch len(c.Args()) {
				case 0: // ex.) unknown
					u, err := handler.get()
					for _, v := range u {
						fmt.Println(v.Word)
					}
					if err != nil {
						return err
					}
				case 1: // ex.) unknown xxx
					var u models.Unknown
					u.Word = c.Args()[0]
					err := handler.store(&u)
					if err != nil {
						return err
					}
				case 2: // ex.) --delete xxx -d
					if c.Bool("d") {
						var u models.Unknown
						u.Word = c.Args()[0]
						err := handler.delete(&u)
						if err != nil {
							return err
						}
						return nil
					} // ex.) --update xxx yyy -u or --update xxx -u yyy
				case 3:
					if c.Bool("u") {
						var u1, u2 models.Unknown
						u1.Word = c.Args()[0]
						u2.Word = c.Args()[2]
						err := handler.update(&u1, &u2)
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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
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

func (h *CliUnknownHandler) update(u1 *models.Unknown, u2 *models.Unknown) error {
	err := h.ukUse.Update(u1, u2)
	if err != nil {
		return err
	}

	return nil
}

func (h *CliUnknownHandler) delete(u *models.Unknown) error {
	err := h.ukUse.Delete(u)
	if err != nil {
		return err
	}

	return nil
}
