package main

import (
	"github.com/nicktming/mydocker/command"
	"github.com/urfave/cli"
	"log"
	"os"
	_ "github.com/nicktming/mydocker/nsenter"
)

func main()  {
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = "implementation of mydocker"

	app.Commands = []cli.Command{
		command.RunCommand,
		command.InitCommand,
		command.CommitCommand,
		command.ListCommand,
		command.LogCommand,
		command.ExecCommand,
		command.StopCommand,
		command.RemoveCommand,
		command.CopyCommand,
		command.NetworkCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
