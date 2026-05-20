package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	slogecho "github.com/samber/slog-echo"

	"ccassociation/internal/config"
	"ccassociation/internal/database"
	"ccassociation/internal/handler"
	"ccassociation/internal/middleware"
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
	lc := net.ListenConfig{}
	ln, err := lc.Listen(ctx, "tcp", addr)
	if err != nil {
		slog.Error("port unavailable", "port", cfg.Port, "error", err)
		os.Exit(1)
	}
	e.Listener = ln

	middleware.Setup(e, cfg)

	// Structured request logger. Registered before route handlers so the
	// middleware chain wraps the response writer first and observes the
	// final status code (including 404s emitted by Echo's HTTPErrorHandler).
	// The previous chi.Logger wrapped via echo.WrapMiddleware missed those
	// writes and logged them as status 000.
	e.Use(slogecho.NewWithConfig(slog.Default(), slogecho.Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,
		WithUserAgent:    true,
		WithRequestID:    true,
	}))

	h := handler.New(cfg, db)
	h.RegisterRoutes(e)

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
