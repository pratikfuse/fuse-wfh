package lib

import (
	"github.com/urfave/cli/v2"
	"log"
)

func LoginUser(ctx *cli.Context) error{

	log.Print("Logged in")
	return nil
}