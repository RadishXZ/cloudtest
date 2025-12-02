package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCloudTestGOCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use : "fg-apiserver",
		Short : "A very lightweiight full go project",
		Long : `A very lightweight full go project, designed to help beginners quickly
		learn Go project developement.`,
		Silenceusage: true,
		RunE: func(cmd *cobra.Command, arg []string) error {
			fmt.Println("Hello CloudTest!")
			return nil
		},
		Args: cobra.NoArgs,
	}
	return cmd
}