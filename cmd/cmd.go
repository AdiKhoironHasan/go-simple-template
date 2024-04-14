package cmd

import (
	"context"
	"go-simple-template/cmd/http"
	"go-simple-template/config"
	"go-simple-template/pkg/tracer"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	config.LoadEnv(".env")
}

func Start() {
	rootCmd := &cobra.Command{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tp, err := tracer.JaegerTraceProvider()
	if err != nil {
		log.Fatalf("failed to create trace provider: %v", err)
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
		log.Fatalf("failed to execute command: %v", err)
	}
}
