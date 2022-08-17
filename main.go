package main

import (
	"log"
	"os"

	"github.com/tesarwijaya/night-owl/cmd"
)

func main() {
	appCmd := cmd.NewCmd()

	if err := appCmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
