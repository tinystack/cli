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

	helloCmd(app)
}

func helloCmd(app *cli.App) {
	var (
		cmdName = "hello"
		cmdDesc = "this is hello command"
		flags   []cli.Flag
		handler cli.Handler
	)

	flags = append(flags, &cli.StringFlag{
		Name:         "name",
		Desc:         "this is name flag",
		Aliases:      []string{"n"},
		DefaultValue: "hello world",
	})

	handler = func(ctx *cli.Context) error {
		fmt.Println("flag name value is:", ctx.String("name"))

		return nil
	}

	app.Command(
		cli.NewCommand(app, cmdName, cmdDesc, flags, handler),
	)
}

func main() {
	if err := app.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
