package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tinystack/cli"
)

var app = cli.NewApp()

func init() {
	app.SetVersion("ver-0.0.1")
	app.SetOutput(os.Stdout)

	bindCmd(app)
}

func bindCmd(app *cli.App) {
	rootCmd := cli.NewCommand(
		app,
		"root",
		"this is root command",
		nil,
		nil,
	)

	sub1Cmd := cli.NewCommand(
		app,
		"sub1",
		"this is sub1 command",
		[]cli.Flag{
			&cli.StringFlag{
				Name:         "name",
				Desc:         "set sub1 command's name",
				Aliases:      []string{"n"},
				DefaultValue: "subDefaultNameValue",
			},
		},
		func(ctx *cli.Context) error {
			fmt.Println("this is sub1 command")
			return nil
		},
	)

	sub2Cmd := cli.NewCommand(
		app,
		"sub2",
		"this is sub2 command",
		[]cli.Flag{
			&cli.StringFlag{
				Name:         "name",
				Desc:         "set sub2 command's name",
				Aliases:      []string{"n"},
				DefaultValue: "subDefaultNameValue",
			},
		},
		func(ctx *cli.Context) error {
			fmt.Println("this is sub2 command")
			return nil
		},
	)

	sub2sonCmd := cli.NewCommand(
		app,
		"sub2son",
		"this is sub2son command",
		[]cli.Flag{
			&cli.StringFlag{
				Name:         "name",
				Desc:         "set sub2son command's name",
				Aliases:      []string{"n"},
				DefaultValue: "subDefaultNameValue",
			},
		},
		func(ctx *cli.Context) error {
			fmt.Println("this is sub2son command")
			return nil
		},
	)
	sub2Cmd.SubCommand(sub2sonCmd)

	rootCmd.SubCommand(sub1Cmd)
	rootCmd.SubCommand(sub2Cmd)

	app.Command(rootCmd)
}

func main() {
	if err := app.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
