package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCloudTestGoCommand() *cobra.Command{
	cmd := &cobra.Command{
		Use: "fg-apiserver",
		Short: "zls cloud app for test",
		Long: "zls cloud app for test, Git:github.com/RadishXZ/cloudtest",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Hello CloudTest!")
			return nil
		},
		Args: cobra.NoArgs,
	}

	return cmd
}