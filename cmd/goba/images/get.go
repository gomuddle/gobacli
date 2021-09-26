package images

import (
	"fmt"
	"github.com/gomuddle/goba"
	"github.com/gomuddle/gobacli/cmd/goba/survey"
	"github.com/gomuddle/gobaclient"
	"github.com/urfave/cli"
	"net/url"
)

// getCommand returns a command for retrieving database images.
func getCommand() cli.Command {
	return cli.Command{
		Name: "get",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "image-name"},
		},
		Before: survey.RequireMixedFlagsFunc(requiredFlags, "image-name"),
		Action: get(),
	}
}

// get sends a request to retrieve the image with the
// given parameters to the server at the specified url.
func get() cli.ActionFunc {
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

		_, err = gobaclient.GetImage(*uri, creds, typ, name)
		if err == nil {
			fmt.Println("Image found")
		}
		return err
	}
}

// getAllCommand returns a command for retrieving a collection of database images.
func getAllCommand() cli.Command {
	return cli.Command{
		Name:   "getall",
		Before: survey.RequireFlagsFunc(requiredFlags...),
		Action: getAll(),
	}
}

// getAll sends a request to retrieve a collection of images
// with the given parameters to the server at the specified
// url.
func getAll() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		creds := gobaclient.Credentials{
			Username: ctx.Parent().Parent().GlobalString("username"),
			Password: ctx.Parent().Parent().GlobalString("password"),
		}

		uri, err := url.Parse(ctx.Parent().Parent().GlobalString("uri"))
		if err != nil {
			return err
		}

		typ := goba.DatabaseType(ctx.Parent().Parent().GlobalString("database-type"))

		images, err := gobaclient.GetAllImages(*uri, creds, typ)
		if err == nil {
			printImageNames(images...)
		}
		return err
	}
}
