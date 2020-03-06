package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// NewRootCmd creates the root collector command
func NewRootCmd() *cobra.Command {
	ctx := context.Background()

	cmd := &cobra.Command{
		Use:   "collector",
		Short: "Collector is the meteo kit utility that collects data published by sensors and propagates it to metrics storage",
	}

	cmd.AddCommand(NewVersionCmd(ctx))
	cmd.AddCommand(NewConsumeCmd(ctx))

	return cmd
}
