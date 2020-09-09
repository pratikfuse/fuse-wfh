package app

import (
	. "github.com/pratikfuse/fuse-wfh/app/lib"
	"github.com/urfave/cli/v2"
)




type App struct {
	cliApp *cli.App
}


func Init() *App{

	app := &App{
		cliApp: &cli.App{
			Name:                   "Fuse WFH",
			HelpName:               "",
			Usage:                  "",
			UsageText:              "",
			ArgsUsage:              "",
			Version:                "",
			Description:            "Work from home cli for Fuse Humans, because some things are more important",
			Commands: GetCommands(),
		},
	}
	return app
}


func (app *App) Run (cliArgs []string) error {
	err := app.cliApp.Run(cliArgs)
	return err
}


