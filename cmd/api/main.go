package main

import (
	"context"
	"time"

	"github.com/anvidev/rma-portal/internal/apidoc"
	"github.com/anvidev/rma-portal/internal/auth"
	"github.com/anvidev/rma-portal/internal/database"
	"github.com/anvidev/rma-portal/internal/env"
	"github.com/anvidev/rma-portal/internal/mailer"
	"github.com/anvidev/rma-portal/internal/queue"
	"github.com/anvidev/rma-portal/internal/store"
	"go.uber.org/zap"
)

func main() {
	config := config{
		server: serverConfig{
			env:          env.String("SERVER_ENV", "development"),
			addr:         env.String("SERVER_ADDR", ":9090"),
			readTimeout:  env.Duration("SERVER_READ_TIMEOUT", time.Second*10),
			writeTimeout: env.Duration("SERVER_WRITE_TIMEOUT", time.Second*30),
			idleTimeout:  env.Duration("SERVER_IDLE_TIMEOUT", time.Minute),
		},
		database: databaseConfig{
			addr:         env.String("POSTGRES_ADDR", "postgres://admin:adminpassword@localhost/rma-db?sslmode=disable"),
			maxOpenConns: env.Int("POSTGRES_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.Int("POSTGRES_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.Duration("POSTGRES_MAX_IDLE_TIME", time.Minute*15),
		},
		auth: authConfig{
			token: tokenConfig{
				host:   env.String("AUTH_TOKEN_HOST", "rma-portal"),
				secret: env.String("AUTH_TOKEN_SECRET", "supersecret"),
				expiry: env.Duration("AUTH_TOKEN_EXPIRY", time.Hour*24*3),
			},
		},
		resend: resendMailerConfig{
			apikey:       env.MustString("RESEND_API_KEY"),
			from:         env.String("RESEND_FROM", "Skancode RMA <noreply@nemunivers.app>"),
			serviceEmail: env.String("RESEND_SERVICE_EMAIL", "av@skancode.dk"),
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

	mail := mailer.NewResendMailer(
		config.resend.apikey,
		config.resend.from,
	)

	queue := queue.New(
		queue.WithEvent(queue.TicketCreated, func(ctx context.Context, payload store.Ticket) error {
			if err := mail.Send([]string{config.resend.serviceEmail}, mailer.TicketCreatedSkancode, payload); err != nil {
				return err
			}
			return nil
		}),
	)

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
		mailer:        mail,
		queue:         queue,
		documentation: documentation,
	}

	queue.Start(context.Background())
	api.addDocs()
	mux := api.mount()

	logger.Fatal(api.run(mux))
}
