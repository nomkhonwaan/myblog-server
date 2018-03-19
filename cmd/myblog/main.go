package main

import (
	"fmt"
	"os"

	"github.com/nomkhonwaan/myblog-server/cmd/myblog/app"
)

func main() {
	cmd := app.New()

	if err := cmd.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "myblog: %v", err)
		os.Exit(1)
	}
}
