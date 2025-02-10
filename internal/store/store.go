package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

const (
	queryTimeoutDuration = time.Second * 5
)

var (
	ErrUserDublicateEmail = errors.New("email already in use")
	ErrUserNotFound       = errors.New("user does not exist")
)

type TODO any

type Store struct {
	Users   userStorer
	Tickets ticketStorer
	Logs    logStorer
}

func NewStore(db *sql.DB) Store {
	return Store{
		Users:   &userStore{db},
		Tickets: &ticketStore{db},
		Logs:    &logStore{db},
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
	GetByID(ctx context.Context, ID int64) (*Ticket, error)
	List(ctx context.Context, filters TicketFilters) ([]Ticket, int, error)
	DeleteByID(ctx context.Context, ID int64) error
}

type logStorer interface {
	Create(ctx context.Context, l *Log) error
	ListByID(ctx context.Context, TicketID int64) ([]Log, error)
}
