package main

import (
	"fmt"
	"os"

	"github.com/pratikfuse/fuse-wfh/app"
)

func main() {
	wfhApp := app.Init()
	if err := wfhApp.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
