package lib

import (
	"github.com/urfave/cli/v2"
)

func GetCommands() []*cli.Command {
	return []*cli.Command {
		{
			Name:  "login",
			Usage: "Login to google suite account",
			Action: Authenticate,
		},
		{
			Name: "check-in",
			Usage: "Check in time",
			Action: CheckIn,
		},
	}
}