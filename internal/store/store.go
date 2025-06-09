package store

import (
	"context"
	"database/sql"
	"time"
)

const (
	queryTimeoutDuration = time.Second * 5
)

type Store struct {
	Users   userStorer
	Tickets ticketStorer
}

func NewStore(db *sql.DB) Store {
	return Store{
		Users:   &userStore{db},
		Tickets: &ticketStore{db},
	}
}

type userStorer interface {
	Create(ctx context.Context, u *User) error
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, ID int64) (User, error)
	List(ctx context.Context) ([]User, error)
}

type ticketStorer interface {
	Create(ctx context.Context, t *Ticket) error
	GetByID(ctx context.Context, ID string) (*Ticket, error)
	List(ctx context.Context, filters TicketFilters) ([]Ticket, int, error)
	DeleteByID(ctx context.Context, ID string) error
	CreateLog(ctx context.Context, l *Log) error
	UpdateLog(ctx context.Context, logID int64, l *Log) error
	ListInternalLogs(ctx context.Context, ID string) ([]Log, error)
	ListExternalLogs(ctx context.Context, ID string) ([]Log, error)
	CreateFile(ctx context.Context, file *File) error
}

func withTx(ctx context.Context, db *sql.DB, fn func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
