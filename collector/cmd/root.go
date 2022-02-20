package cmd

import "github.com/spf13/cobra"

var version = "0.0.0-dev"

// NewRootCmd creates the root collector command
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "collector",
		Short:   "Collector is the meteo kit utility that collects data published by sensors and propagates it to metrics storage",
		Version: version,
	}

	cmd.AddCommand(NewConsumeCmd())

	return cmd
}
