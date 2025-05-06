package main

import (
	"time"

	"github.com/anvidev/rma-portal/internal/apidoc"
	"github.com/anvidev/rma-portal/internal/auth"
	"github.com/anvidev/rma-portal/internal/database"
	"github.com/anvidev/rma-portal/internal/env"
	"github.com/anvidev/rma-portal/internal/store"
	"go.uber.org/zap"
)

func main() {
	config := config{
		server: serverConfig{
			env:          env.GetString("SERVER_ENV", "development"),
			addr:         env.GetString("SERVER_ADDR", ":8080"),
			readTimeout:  env.GetDuration("SERVER_READ_TIMEOUT", time.Second*10),
			writeTimeout: env.GetDuration("SERVER_WRITE_TIMEOUT", time.Second*30),
			idleTimeout:  env.GetDuration("SERVER_IDLE_TIMEOUT", time.Minute),
		},
		database: databaseConfig{
			addr:         env.GetString("POSTGRES_ADDR", "postgres://admin:adminpassword@localhost/rma-db?sslmode=disable"),
			maxOpenConns: env.GetInt("POSTGRES_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("POSTGRES_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetDuration("POSTGRES_MAX_IDLE_TIME", time.Minute*15),
		},
		auth: authConfig{
			token: tokenConfig{
				host:   env.GetString("AUTH_TOKEN_HOST", "rma-portal"),
				secret: env.GetString("AUTH_TOKEN_SECRET", "supersecret"),
				expiry: env.GetDuration("AUTH_TOKEN_EXPIRY", time.Hour*24*3),
			},
		},
	}

	documentation := apidoc.NewDocumentation(apidoc.Info{
		Title:       "RMA Portal",
		Description: "Backend API for Skancode A/S RMA service portal",
		Version:     "0.0.1",
		Contact: apidoc.InfoContact{
			Name:  "anvi",
			Email: "hello@anvi.dev",
		},
	})

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := database.NewPostgreSQL(
		config.database.addr,
		config.database.maxOpenConns,
		config.database.maxIdleConns,
		config.database.maxIdleTime,
	)
	if err != nil {
		logger.Fatalw("database connection failed", "error", err)
	}

	store := store.NewStore(db)
	authenticator := auth.NewJWTAuthenticator(
		config.auth.token.secret,
		config.auth.token.host,
		config.auth.token.host,
	)

	api := &api{
		logger:        logger,
		config:        config,
		store:         store,
		auth:          authenticator,
		documentation: documentation,
	}

	api.addDocs()
	mux := api.mount()

	logger.Fatal(api.run(mux))
}
