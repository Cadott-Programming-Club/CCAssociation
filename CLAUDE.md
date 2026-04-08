# Cadott Community Association Website

Go + Templ + HTMX + Tailwind CSS v4 web application for the Cadott Community Association.

## Critical: Development Rules

- The user is ALWAYS running `make dev` — NEVER start, build, or run the server yourself
- Check `./tmp/air-combined.log` for build errors after making code changes
- NEVER run `go build`, `go run`, or start the server — it's already running
- To verify changes, check the logs or use Chrome DevTools MCP to view the running site
- Old static HTML files are still in the project root — open them directly as `file://` paths for comparison

## Development Workflow

- `make dev` is always running during development
- It automatically: kills existing process, regenerates Templ, runs go mod tidy, rebuilds and restarts
- Developer does NOT need to manually run: templ generate, go build, air
- Run `make css-watch` in a separate terminal for Tailwind CSS hot reload

## Environment

- All config via `.envrc` with direnv (`direnv allow`)
- DATABASE_URL (SQLite: `./data/ccassociation.db`)
- PORT (default 8000), ENV, LOG_LEVEL

## Key Commands

| Command | Purpose |
|---------|---------|
| `make dev` | Start with hot reload (main workflow) |
| `make build` | Build production binary |
| `make test` | Run tests with race detection |
| `make lint` | Run linters |
| `make migrate` | Run database migrations |
| `make css-watch` | Watch Tailwind (separate terminal) |
| `make setup` | Install dev tools (air, templ, sqlc, goose, golangci-lint) |

## Project Structure

| Directory | Purpose |
|-----------|---------|
| `cmd/server/` | Entry point, slog init, go:generate directives |
| `internal/config/` | Environment config (Config struct) |
| `internal/database/` | SQLite database, embedded migrations |
| `internal/handler/` | HTTP handlers (one file per page) |
| `internal/middleware/` | Echo middleware (CORS, gzip, security, site config) |
| `internal/meta/` | SEO meta tags (PageMeta struct, context helpers) |
| `internal/ctxkeys/` | Typed context keys |
| `templates/layouts/` | Base layout, meta tags, header/footer |
| `templates/pages/` | Page templates (home, events, gallery, faq, contact) |
| `static/` | CSS, JS, images |
| `sqlc/` | SQL queries and sqlc config |

## Code Patterns

- **Logging**: Use `slog` (never `fmt.Printf` or `log.Printf`)
- **Errors**: Wrap with context using `fmt.Errorf`
- **Database**: SQLite via `modernc.org/sqlite` (CGO-free), goose migrations
- **Templates**: Templ components, meta owned by templates not handlers
- **Routing**: Echo v4 with chi middleware logger
- **Deployment**: Dokploy with Nixpacks (`nixpacks.toml`)

## Content

- Images hosted externally on `cadottcommunity.com`
- Content is currently hardcoded in Templ templates
- Facebook embeds used for timeline, events, and photos
- Google Maps embed on contact page
