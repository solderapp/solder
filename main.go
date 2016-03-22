package main

import (
	"os"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/solderapp/solder-api/cmd"
	"github.com/solderapp/solder-api/config"
)

var (
	version string
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cfg := &config.Config{
		Version: version,
	}

	app := cli.NewApp()
	app.Name = "solder"
	app.Version = version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "Manage mod packs for the Technic launcher"

	app.Before = func(c *cli.Context) error {
		logrus.SetOutput(os.Stdout)

		if cfg.Debug {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}

		return nil
	}

	app.Commands = []cli.Command{
		cmd.Server(cfg),
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Show the help, so what you see now",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Print the current version of that tool",
	}

	app.Run(os.Args)
}
