package cmd

import (
	"context"
	"go-simple-template/cmd/http"
	"go-simple-template/config"
	"go-simple-template/pkg/logger"
	"go-simple-template/pkg/tracer"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	config.LoadEnv(".env")
}

func Start() {
	rootCmd := &cobra.Command{}

	log := logger.Init()
	defer log.Sync()

	ctx := logger.WithCtx(context.Background(), log)

	tp, err := tracer.JaegerTraceProvider()
	if err != nil {
		log.Fatal("Failed to create trace provider", zap.Error(err))
	}

	tracer.RegisterTracer(tp)

	cmd := []*cobra.Command{
		{
			Use:   "http",
			Short: "Http is a simple HTTP REST API application",
			Run: func(cmd *cobra.Command, args []string) {
				http.Start(ctx)
			},
		},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Failed to execute command", zap.Error(err))
	}
}
