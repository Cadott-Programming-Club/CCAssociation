package main

import (
	"ccassociation/internal/config"
	"ccassociation/internal/database"
	"ccassociation/internal/handler"
	"ccassociation/internal/middleware"
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()
	db, err := database.New(ctx, cfg.DatabaseURL)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	addr := ":" + cfg.Port
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		slog.Error("port unavailable", "port", cfg.Port, "error", err)
		os.Exit(1)
	}
	e.Listener = ln

	middleware.Setup(e, cfg)

	h := handler.New(cfg, db)
	h.RegisterRoutes(e)

	e.Use(echo.WrapMiddleware(chimw.Logger))

	go func() {
		slog.Info("starting server", "url", fmt.Sprintf("http://localhost:%s", cfg.Port), "env", cfg.Env)
		if err := e.Start(""); err != nil {
			slog.Info("shutting down server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		slog.Error("server shutdown error", "error", err)
	}

	slog.Info("server stopped")
}

