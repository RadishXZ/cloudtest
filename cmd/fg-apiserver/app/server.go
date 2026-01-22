package app

import (
	"github.com/RadishXZ/cloudtest/cmd/fg-apiserver/app/options"
	"github.com/RadishXZ/cloudtest/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

func NewCloudTestGOCommand() *cobra.Command {
	opts := options.NewServerOptions()

	cmd := &cobra.Command{
		Use : "fg-apiserver",
		Short : "A very lightweiight full go project",
		Long : `A very lightweight full go project, designed to help beginners quickly
		learn Go project developement.`,
		SilenceUsage: true,
		
		RunE: func(cmd *cobra.Command, arg []string) error {
			return run(opts)
		},
		Args: cobra.NoArgs,
	}

	cobra.OnInitialize(onInitialize)

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the fg-apiserver configuration file.")
	
	version.AddFlags(cmd.PersistentFlags())
	
	return cmd
}

func run (opts *options.ServerOptions) error {
	if err := viper.Unmarshal(opts); err != nil {
		return err
	}

	if err := opts.Validate(); err != nil {
		return err
	}

	cfg, err := opts.Config()
	if err != nil {
		return err
	}

	server, err := cfg.NewServer()
	if err != nil {
		return err
	}

	version.PrintAndExitIfRequested()

	return server.Run()
}