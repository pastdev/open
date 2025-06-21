package main

import (
	"fmt"
	"os"

	"github.com/pastdev/open/pkg/open"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use:   "open",
		Short: `A tool for opening URL's using their OS defined default application.`,
		//nolint: revive // required to match upstream signature
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := open.Open(args[0])
			if err != nil {
				return fmt.Errorf("main open: %w", err)
			}
			return nil
		},
	}

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
