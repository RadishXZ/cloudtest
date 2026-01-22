package app

import (
	"io"
	"log/slog"
	"os"

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

	initLog()

	return server.Run()
}

func initLog() {
	format := viper.GetString("log.format")
	level := viper.GetString("log.level")
	output := viper.GetString("log.output")

	var slevel slog.Level
	switch level {
	case "debug":
		slevel = slog.LevelDebug
	case "info":
		slevel = slog.LevelInfo
	case "warn":
		slevel = slog.LevelWarn
	case "error":
		slevel = slog.LevelError
	default:
		slevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{Level: slevel}

	var w io.Writer
	var err error
	switch output {
	case "":
		w = os.Stdout
	case "stdout":
		w = os.Stdout
	default:
		w, err = os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
	}

	if err != nil {
		return
	}

	var handler slog.Handler
	switch format {
	case "json":
		handler = slog.NewJSONHandler(w, opts)
	case "text":
		handler = slog.NewTextHandler(w, opts)
	default:
		handler = slog.NewJSONHandler(w, opts)
	}

	slog.SetDefault(slog.New(handler))
}