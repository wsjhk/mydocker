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

	app.Flags = []cli.Flag {
		cli.BoolFlag{
			Name: "flag",
			Usage: "enable tty",
		},
		cli.StringFlag{
			Name: "lang",
			Value: "english",
		},
	}

	app.Action = func(c *cli.Context) error {
		log.Printf("args:%s\n", c.Args())
		log.Printf("flag:%v\n", c.Bool("flag"))
		log.Printf("lang:%s\n", c.String("lang"))
		for i := 0; i < 5; i++ {
			log.Printf("sleep %d\n", i)
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