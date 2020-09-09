package main



import (
	"fmt"
	"github.com/pratikfuse/fuse-wfh/app"
	"os"
)

func main() {
	wfhApp := app.Init()
	if err := wfhApp.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
