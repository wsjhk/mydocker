package command

import (
	"github.com/nicktming/mydocker/cgroups"
	"github.com/nicktming/mydocker/cgroups/subsystems"
	"github.com/urfave/cli"
)

var RunCommand = cli.Command{
	Name: "run",
	Flags: []cli.Flag {
		cli.BoolFlag{
			Name: "it",
			Usage: "enable tty",
		},
		cli.StringFlag{
			Name: "m",
			Usage: "memory usage",
		},
		cli.StringFlag{
			Name: "r",
			Usage: "set rootPath",
		},
	},
	Action: func(c *cli.Context) error {
		tty 	:= c.Bool("it")
		memory  := c.String("m")
		rootPath  := c.String("r")
		command := c.Args().Get(0)

		res := subsystems.ResourceConfig{
			MemoryLimit: memory,
		}
		cg := cgroups.CroupManger {
			Resource: &res,
			SubsystemsIns: make([]subsystems.Subsystem, 0),
		}
		if memory != "" {
			cg.SubsystemsIns = append(cg.SubsystemsIns, &subsystems.MemorySubsystem{})
		}

		Run(command, tty, &cg, rootPath)
		return nil
	},
}

var InitCommand = cli.Command{
	Name: "init",
	Flags: []cli.Flag {
		cli.BoolFlag{
			Name: "it",
			Usage: "enable tty",
		},
	},
	Action: func(c *cli.Context) error {
		command := c.Args().Get(0)
		Init(command)
		return nil
	},
}
