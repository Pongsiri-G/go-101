package database

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/graphzc/go-clean-template/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

// @WireSet("Infrastructure")
func NewSQLXClient(ctx context.Context, config *config.Config) *sqlx.DB {
	// If DATABASE_URI is not provided, attempt to build it from component env vars
	uri := config.Database.URI
	if uri == "" {
		user := os.Getenv("DATABASE_USERNAME")
		pass := os.Getenv("DATABASE_PASSWORD")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")
		name := os.Getenv("DATABASE_NAME")
		ssl := os.Getenv("DATABASE_SSL_MODE")
		if ssl == "" {
			ssl = "disable"
		}
		// Build a postgres URI: postgres://user:pass@host:port/name?sslmode=...
		// URL-escape the password
		uri = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, url.QueryEscape(pass), host, port, name, ssl)
	}

	// Mask password when logging
	maskedURI := uri
	if u, err := url.Parse(uri); err == nil {
		if u.User != nil {
			username := u.User.Username()
			// replace password with ****
			u.User = url.UserPassword(username, "****")
			maskedURI = u.String()
		}
	}

	log.Info().Str("database_uri", maskedURI).Msg("Database connection string (masked)")

	db, err := sqlx.ConnectContext(ctx, config.Database.Driver, uri)
	if err != nil {
		log.Panic().
			Err(err).
			Msg("Failed to connect to the database")
	}

	log.Info().
		Msg("Connected to the database")

	return db
}
