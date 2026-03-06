package main

import (
	"log/slog"

	"github.com/p-alvarenga/go_tcp-tracker-server/internal/config"
	"github.com/p-alvarenga/go_tcp-tracker-server/internal/server"
)

func main() {
	cfg := config.DefaultConfig()

	server := server.NewServer(cfg, slog.Default())
	server.Boot()
}
