package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/urfave/cli"
)

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalln(err)
	}
	time.Local = loc

	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	os.Exit(cliMain())
}

func cliMain() int {
	app := cli.NewApp()
	app.Name = "isupipebench"
	app.Usage = "isupipe ベンチマーカー"
	app.Description = "isupipeのベンチマークを実施"
	app.HelpName = "isupipebench"

	app.Commands = []cli.Command{
		run,
	}

	app.Action = func(cliCtx *cli.Context) error {
		return cli.ShowAppHelp(cliCtx)
	}

	if err := app.Run(os.Args); err != nil {
		exitErr := err.(*cli.ExitError)
		log.Println(exitErr.Error())
		return exitErr.ExitCode()
	}

	return 0
}
