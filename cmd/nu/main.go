package main

import (
	"encoding/json"
	"fmt"
	"github.com/SirMetathyst/NovelUpdatesClient"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:    "nu",
		Version: "v0.1",
		Usage:   "NovelUpdate Client",
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
		Commands: []*cli.Command{
			{
				Name: "search",
				Flags: []cli.Flag{

					// PAGINATION
					&cli.IntFlag{Name: "page", Aliases: []string{"p"}},
					//&cli.IntFlag{Name: "page-from"},
					//&cli.IntFlag{Name: "page-to"},
					//&cli.IntFlag{Name: "items-start"},
					//&cli.IntFlag{Name: "items-end"},

					// JSON
					&cli.BoolFlag{Name: "json"},
					&cli.BoolFlag{Name: "json-indent"},
				},
				Action: func(context *cli.Context) error {
					response, err := NovelUpdatesClient.DoSearchRequestWith(&NovelUpdatesClient.BasicRequester{}, &NovelUpdatesClient.SearchQuery{
						Page: context.Int("page"),
					})
					if err != nil {
						return err
					}
					if context.Bool("json") {
						if !context.Bool("json-indent") {
							v, err := json.Marshal(response)
							if err != nil {
								return err
							}
							fmt.Println(string(v))
						} else {
							v, err := json.MarshalIndent(response, "", "\t")
							if err != nil {
								return err
							}
							fmt.Println(string(v))
						}
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
