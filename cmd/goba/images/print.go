package images

import (
	"fmt"
	"github.com/gomuddle/goba"
)

// printImageNames is a helper function for pretty-printing a list of images.
func printImageNames(images ...goba.Image) {
	if len(images) == 1 {
		fmt.Println("Image:", images[0].Name)
	}
	if len(images) > 1 {
		fmt.Println("Images: [")
		for _, img := range images {
			fmt.Println("\t" + img.Name)
		}
		fmt.Print("]")
	}
}
