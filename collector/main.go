package main

import (
	"log"

	"github.com/vgarvardt/meteo/collector/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
