package testutil

import (
	"ccassociation/internal/config"
	"ccassociation/internal/database"
	"context"
	"testing"
)

func NewTestDB(t *testing.T) *database.DB {
	t.Helper()

	ctx := context.Background()
	db, err := database.New(ctx, ":memory:")
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	return db
}

func NewTestConfig(t *testing.T) *config.Config {
	t.Helper()

	return &config.Config{
		DatabaseURL: ":memory:",
		Port:        "0",
		Env:         "test",
		Site: config.SiteConfig{
			Name: "Cadott Community Association",
			URL:  "http://localhost:8000",
		},
	}
}
