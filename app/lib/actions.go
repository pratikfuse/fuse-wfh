package lib

import (
	"fmt"
	"log"

	"github.com/pratikfuse/fuse-wfh/app/lib/auth"
	sh "github.com/pratikfuse/fuse-wfh/app/lib/sheet"
	"github.com/urfave/cli/v2"
)

var sheet sh.Sheet

func Authenticate(ctx *cli.Context) error {

	port := auth.GetRedirectPort()
	oauthUrl := auth.GetOauthUrl(port)
	err := auth.OpenOauthUrl(oauthUrl)

	if err != nil {
		return err
	}

	server := &auth.Server{}
	errChan := make(chan error)
	shutdownSignal := make(chan bool)
	server.RunAuthServer(port, errChan, shutdownSignal)
	if err := <-errChan; err != nil {
		return err
	}
	log.Println("logged in to fuse wfh")
	return nil
}

func CheckIn(ctx *cli.Context) error {
	//accessCredentials := auth.ParseCredentials()
	s, err := sh.GetInstance("API_KEY")
	if err != nil {
		return err
	}
	sheet := s.Service.Spreadsheets.Get("SHEETS_ID")
	fmt.Print(sheet.Ranges())

	return nil
}
