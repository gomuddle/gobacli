package images

import (
	"fmt"
	"github.com/gomuddle/goba"
	"github.com/gomuddle/gobacli/cmd/goba/survey"
	"github.com/gomuddle/gobaclient"
	"github.com/urfave/cli"
	"net/url"
)

// deleteCommand returns a command for deleting database images.
func deleteCommand() cli.Command {
	return cli.Command{
		Name: "delete",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "image-name"},
		},
		Before: survey.RequireMixedFlagsFunc(requiredFlags, "image-name"),
		Action: delete(),
	}
}

// delete sends a request to delete the image with the
// given parameters to the server at the specified url.
func delete() cli.ActionFunc {
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

		if err = gobaclient.DeleteImage(*uri, creds, typ, name); err != nil {
			return err
		}

		fmt.Println("Image deleted")
		return nil
	}
}
