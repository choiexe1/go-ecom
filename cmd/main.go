package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/choiexe1/go-ecom/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	config := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DB_STRING", "host=localhost user=postgres password=postgres dbname=ecom_db sslmode=disable"),
		},
	}

	// logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := pgx.Connect(
		ctx,
		config.db.dsn,
	)

	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("Connected to Database", "dsn", config.db.dsn)

	api := application{
		config: config,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server has failed to start", "error", err)
		os.Exit(1)
	}
}
