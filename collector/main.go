package main

import (
	"context"
	"log"

	"github.com/vgarvardt/meteo/collector/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	ctx := context.Background()
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
