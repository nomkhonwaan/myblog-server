package main

import (
	"os"

	"github.com/nomkhonwaan/myblog-server/cmd/myblog/app"
	"github.com/sirupsen/logrus"
)

func main() {
	cmd := app.New()
	err := cmd.Run(os.Args)

	if err != nil {
		logrus.Fatalf("myblog: %v", err)
	}
}
