package store

import (
	"context"
	"database/sql"
)

type Log struct {
	ID              int64  `json:"id"`
	TicketID        int64  `json:"ticket_id,omitempty"`
	Status          Status `json:"status"`
	Initiator       string `json:"initiator"`
	ExternalComment string `json:"external_comment"`
	InternalComment string `json:"internal_comment"`
	Inserted        string `json:"inserted"`
}

type logStore struct {
	db *sql.DB
}

func (s *logStore) Create(ctx context.Context, l *Log) error {
	stmt := `
		INSERT INTO logs (ticket_id, status, initiator, external_comment, internal_comment)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, inserted
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		stmt,
		l.TicketID,
		l.Status,
		l.Initiator,
		l.ExternalComment,
		l.InternalComment,
	).Scan(
		&l.ID,
		&l.Inserted,
	)
	if err != nil {
		switch {
		case err.Error() == `pq: insert or update on table "logs" violates foreign key constraint "logs_tickets_id_fkey"`:
			return ErrTicketNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *logStore) ListByID(ctx context.Context, ticketID int64) ([]Log, error) {
	stmt := `
		SELECT id, status, initiator, external_comment, internal_comment, inserted
		FROM logs
		WHERE ticket_id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	logs := []Log{}

	rows, err := s.db.QueryContext(
		ctx,
		stmt,
		ticketID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var l Log

		if err := rows.Scan(
			&l.ID,
			&l.Status,
			&l.Initiator,
			&l.ExternalComment,
			&l.InternalComment,
			&l.Inserted,
		); err != nil {
			return nil, err
		}

		logs = append(logs, l)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}
