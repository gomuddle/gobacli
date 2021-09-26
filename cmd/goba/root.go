package goba

import (
	"github.com/gomuddle/gobacli/cmd/goba/images"
	"github.com/urfave/cli"
	"os"
)

// App is the entry point of the cli.
func App() *cli.App {
	return &cli.App{
		Name:  "goba",
		Flags: images.RequiredFlags(),
		Commands: []cli.Command{
			images.RootCommand(),
		},
		BashComplete: cli.DefaultAppComplete,
		Writer:       os.Stdout,
	}
}
