package main

import (
	"github.com/nicktming/mydocker/command"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main()  {
	log.Printf("Main start!\n")
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = "implementation of mydocker"

	app.Commands = []cli.Command{
		command.RunCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Main end!\n")
}
