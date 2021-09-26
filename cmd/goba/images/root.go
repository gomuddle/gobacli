package images

import (
	"github.com/urfave/cli"
)

var requiredFlags = []string{
	"uri",
	"username",
	"password",
	"database-type",
}

// RootCommand returns a command containing subcommands for
// interacting with a server that manages database images.
func RootCommand() cli.Command {
	return cli.Command{
		Name: "images",
		Subcommands: []cli.Command{
			getCommand(),
			getAllCommand(),
			createCommand(),
			applyCommand(),
			deleteCommand(),
		},
		BashComplete: cli.DefaultAppComplete,
	}
}

// RequiredFlags is a helper function that constructs
// a slice of flags from the requiredFlags.
func RequiredFlags() (flags []cli.Flag) {
	for _, flag := range requiredFlags {
		flags = append(flags, cli.StringFlag{Name: flag})
	}
	return
}
