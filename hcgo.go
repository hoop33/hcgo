package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/hoop33/hcgo/model"
)

func main() {
	app := cli.NewApp()
	app.Name = "hcgo"
	app.Version = "0.1.0"
	app.Usage = "Display color information"
	app.Authors = []cli.Author{
		{
			Name:  "Rob Warner",
			Email: "rwarner@grailbox.com",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "show",
			Usage: "Shows a color",
			Action: func(c *cli.Context) {
				color, err := model.NewColorFromHexCode(c.Args().First())
				if err == nil {
					fmt.Printf("%s", color.String())
				} else {
					fmt.Println(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
