package images

import (
	"github.com/gomuddle/goba"
	"github.com/gomuddle/gobacli/cmd/goba/survey"
	"github.com/gomuddle/gobaclient"
	"github.com/urfave/cli"
	"net/url"
)

// createCommand returns a command for creating database images.
func createCommand() cli.Command {
	return cli.Command{
		Name:   "create",
		Before: survey.RequireGlobalFlagsFunc(requiredFlags...),
		Action: create(),
	}
}

// create sends a requests for creating a new image with
// the given parameters to the server at the specified url.
func create() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		creds := gobaclient.Credentials{
			Username: ctx.Parent().GlobalString("username"),
			Password: ctx.Parent().GlobalString("password"),
		}

		uri, err := url.Parse(ctx.Parent().GlobalString("uri"))
		if err != nil {
			return err
		}

		typ := goba.DatabaseType(ctx.Parent().GlobalString("database-type"))

		image, err := gobaclient.CreateImage(*uri, creds, typ)
		if err == nil {
			printImageNames(*image)
		}
		return err
	}
}
