package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/LeandroDeJesus-S/rinha25/internal/config"
	"github.com/redis/go-redis/v9"
)

func main() {
	conf := config.New()

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		
	})
	defer rdb.Close()

	if err := rdb.Ping(ctx).Err(); err != nil {
		slog.Error("Failed to connect to Redis", "err", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	appAddr := fmt.Sprintf(":%d", conf.AppPort)
	slog.Info("Starting API server", "addr", appAddr)
	if err := http.ListenAndServe(appAddr, mux); err != nil {
		slog.Error("Failed to start API server", "err", err)
		os.Exit(1)
	}
}
