package main

import (
	"github.com/gomuddle/gobacli/cmd/goba"
	"log"
	"os"
)

func main() {
	if err := goba.App().Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
