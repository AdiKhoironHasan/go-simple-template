package cmd

import (
	"context"
	"go-simple-template/cmd/api/rest"
	"go-simple-template/cmd/consumer"
	"go-simple-template/cmd/migration"
	"go-simple-template/config"
	"go-simple-template/pkg/logger"
	"go-simple-template/pkg/tracer"

	cons "go-simple-template/internal/consumer"

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

	restCmd := cobra.Command{
		Use:   "rest",
		Short: "rest is a simple REST API application",
		Run: func(cmd *cobra.Command, args []string) {
			rest.Start(ctx)
		},
	}

	consumerCmd := cobra.Command{
		Use:   "consumer",
		Short: "consumer is a simple consumer application",
	}

	consumerPingCmd := cobra.Command{
		Use:   "ping",
		Short: "ping is a simple consumer application",
		Run: func(cmd *cobra.Command, args []string) {
			consumer.Start(ctx, cons.NewPing())
		},
	}

	migrationCmd := cobra.Command{
		Use:   "migration",
		Short: "migration is for running database migration",
	}

	migrationAutoMigrateCmd := cobra.Command{
		Use:   "automigrate",
		Short: "automigrate is for running database automigrate",
		Run: func(cmd *cobra.Command, args []string) {
			migration.AutoMigrate()
			log.Info("Automigrate success")
		},
	}

	// Register command to consumer command
	consumerCmd.AddCommand(&consumerPingCmd)

	// Register command to migration command
	migrationCmd.AddCommand(&migrationAutoMigrateCmd)

	// Register command to root command
	rootCmd.AddCommand(
		&restCmd,
		&consumerCmd,
		&migrationCmd,
	)

	// Execute root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Failed to execute command", zap.Error(err))
	}
}
