package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	log.Printf("Main start!\n")
	app := cli.NewApp()
	app.Name = "example"
	app.Usage = "make an explosive entrance"

	runCommand := cli.Command{
		Name: "run",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name: "it",
				Usage: "enable tty",
			},
		},
		Action: func(c *cli.Context) error {
			log.Printf("runcommand args:%s\n", c.Args())
			log.Printf("runcommand tty:%v\n", c.Bool("it"))
			for i := 0; i < 5; i++ {
				log.Printf("runcommand sleep %d\n", i)
				time.Sleep(1 * time.Second)
			}
			return nil
		},
	}

	app.Commands = []cli.Command {
		runCommand,
	}

	app.Action = func(c *cli.Context) error {
		log.Printf("main function args:%s\n", c.Args())
		for i := 0; i < 5; i++ {
			log.Printf("main function sleep %d\n", i)
			time.Sleep(1 * time.Second)
		}
		return nil
	}

	log.Printf("Before invoking Run!\n")
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Main end!\n")
}