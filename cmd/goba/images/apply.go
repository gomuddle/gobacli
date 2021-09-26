package images

import (
	"github.com/gomuddle/goba"
	"github.com/gomuddle/gobacli/cmd/goba/survey"
	"github.com/gomuddle/gobaclient"
	"github.com/urfave/cli"
	"net/url"
)

// applyCommand returns a command for applying database images to databases.
func applyCommand() cli.Command {
	return cli.Command{
		Name: "apply",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "image-name"},
		},
		Before: survey.RequireMixedFlagsFunc(requiredFlags, "image-name"),
		Action: apply(),
	}
}

// apply sends a request for applying the image with the given
// parameters to the database with the given type to the server
// at the specified url.
func apply() cli.ActionFunc {
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
		name := ctx.String("image-name")

		return gobaclient.ApplyImage(*uri, creds, typ, name)
	}
}
