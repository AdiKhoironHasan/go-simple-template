package main

import (
	"context"
	"log"
	"log/slog"

	"go-simple-template/cmd/api/rest"
	"go-simple-template/internal/infrastructure"
	"go-simple-template/internal/pkg/config"

	"github.com/spf13/cobra"
)

func main() {
	var (
		ctx      = context.Background()
		rootCmd  = cobra.Command{}
		logLevel = slog.LevelInfo
	)

	config.LoadEnv(".env")

	if config.AppDebug() {
		logLevel = slog.LevelDebug
	}

	infrastructure.NewSlog(logLevel)

	cmd := []*cobra.Command{
		{
			Use:   "rest",
			Short: "Start the REST API server",
			Run: func(cmd *cobra.Command, args []string) {
				rest.Start(ctx)
			},
		},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		log.Fatalf("failed to execute command: %v", err)
	}
}
