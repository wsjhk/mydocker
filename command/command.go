package command

import (
	"github.com/nicktming/mydocker/cgroups"
	"github.com/nicktming/mydocker/cgroups/subsystems"
	"github.com/urfave/cli"
	"log"
	"os"
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
		cli.StringSliceFlag{
			Name: "v",
			Usage: "enable volume",
		},
		cli.BoolFlag{
			Name: "d",
			Usage: "enable detach",
		},
		cli.StringFlag{
			Name: "name",
			Usage: "container name",
		},
		/*
		cli.StringFlag{
			Name: "v",
			Usage: "enable volume",
		},
		*/
	},
	Action: func(c *cli.Context) error {
		tty 	  := c.Bool("it")
		memory    := c.String("m")
		rootPath  := c.String("r")
		//volume    := c.String("v")
		volumes   := c.StringSlice("v")
		detach    := c.Bool("d")
		containerName    := c.String("name")
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

		if detach {
			tty = false
		}

		Run(command, tty, &cg, rootPath, volumes, containerName)
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

var CommitCommand = cli.Command{
	Name: "commit",
	Action: func(c *cli.Context) error {
		imageName := c.Args().Get(0)
		Commit(imageName)
		return nil
	},
}

var ListCommand = cli.Command{
	Name: "ps",
	Action: func(c *cli.Context) error {
		List()
		return nil
	},
}

var LogCommand = cli.Command{
	Name: "logs",
	Action: func(c *cli.Context) error {
		containerName := c.Args().Get(0)
		Logs(containerName)
		return nil
	},
}

var ExecCommand = cli.Command{
	Name: "exec",
	Action: func(c *cli.Context) error {
		if os.Getenv("mydocker_pid") != "" {
			log.Printf("pid callback pid %s", os.Getgid())
			return nil
		}
		containerName := c.Args().Get(0)
		command 	  := c.Args().Get(1)
		log.Printf("containerName:%s,command:%s\n", containerName, command)
		Exec(containerName, command)
		return nil
	},
}

var StopCommand = cli.Command{
	Name: "stop",
	Usage: "stop a container",
	Action: func(c *cli.Context) error {
		containerName := c.Args().Get(0)
		Stop(containerName)
		return nil
	},
}

var RemoveCommand = cli.Command{
	Name: "rm",
	Usage: "remove a stopped container",
	Action: func(c *cli.Context) error {
		containerName := c.Args().Get(0)
		Remove(containerName)
		return nil
	},
}

