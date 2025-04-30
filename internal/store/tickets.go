package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
)

type Status int
type Category int

const (
	OPEN Status = iota + 1
	CLOSED
	EXT_REPAIR
	INT_REPAIR
	QUOTE_SENT
	QUOTE_ACCEPT
	QOUTE_REJECT_RETURN
	QUOTE_REJECT_DESTROY
)

const (
	SOFTWARE Category = iota + 1
	HARDWARE
)

var (
	ErrStatusNotImplemented   = fmt.Errorf("status ikke implementeret")
	ErrCategoryNotImplemented = fmt.Errorf("kategori ikke implementeret")
	ErrTicketNotFound         = fmt.Errorf("sag blev ikke fundet")
)

var Statuses = map[Status]string{
	OPEN:                 "åben",
	CLOSED:               "lukket",
	EXT_REPAIR:           "ekstern reperation",
	INT_REPAIR:           "intern reperation",
	QUOTE_SENT:           "tilbud sendt",
	QUOTE_ACCEPT:         "tilbud accepteret",
	QOUTE_REJECT_RETURN:  "tilbud afvist (sendes tilbage)",
	QUOTE_REJECT_DESTROY: "tilbud afvist (kasseres)",
}

var Categories = map[Category]string{
	SOFTWARE: "software",
	HARDWARE: "hardware",
}

func (s Status) MarshalJSON() ([]byte, error) {
	str, ok := Statuses[s]
	if !ok {
		return nil, ErrStatusNotImplemented
	}
	return json.Marshal(str)
}

func (s *Status) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	for k, v := range Statuses {
		if str == v {
			*s = k
			return nil
		}
	}
	return ErrStatusNotImplemented
}

func (c Category) MarshalJSON() ([]byte, error) {
	str, ok := Categories[c]
	if !ok {
		return nil, ErrCategoryNotImplemented
	}
	return json.Marshal(str)
}

func (c *Category) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	for k, v := range Categories {
		if str == v {
			*c = k
			return nil
		}
	}
	return ErrCategoryNotImplemented
}

func (c *Category) Scan(value any) error {
	switch v := value.(type) {
	case []byte:
		src := string(value.([]byte))
		catInt, err := strconv.Atoi(src)
		if err != nil {
			return err
		}
		for k := range Categories {
			if Category(catInt) == k {
				*c = Category(catInt)
			}
		}
	default:
		return fmt.Errorf("kategori ikke implementeret: %T", v)
	}
	return nil
}

type TicketFilters struct {
	Query      string      `json:"query"`
	Status     Status      `json:"status"`
	Categories []Category  `json:"categories"`
	Inserted   []time.Time `json:"inserted"`
	Updated    []time.Time `json:"updated"`
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
}

func (f *TicketFilters) Parse(r *http.Request) error {
	q := r.URL.Query()

	f.Limit = 50
	f.Offset = 0

	if q.Has("limit") {
		max := 1000
		int, err := strconv.Atoi(q.Get("limit"))
		if err != nil {
			return err
		}
		if int > max {
			return fmt.Errorf("limit grænse er overskredet, det maksimale antal på limit er %d", max)
		}
		f.Limit = int
	}

	if q.Has("offset") {
		int, err := strconv.Atoi(q.Get("offset"))
		if err != nil {
			return err
		}
		f.Offset = int
	}

	if q.Has("query") {
		f.Query = q.Get("query")
	}

	if q.Has("status") {
		statusQuery := q.Get("status")
		statusFound := false
		for status, label := range Statuses {
			if label == statusQuery {
				f.Status = status
				statusFound = true
				break
			}
		}
		if !statusFound {
			return ErrStatusNotImplemented
		}
	}

	if q.Has("categories") {
		categoryQuery := strings.Split(q.Get("categories"), ",")
		for _, query := range categoryQuery {
			if query == "" {
				break
			}
			categoryFound := false
			for category, label := range Categories {
				if label == query {
					f.Categories = append(f.Categories, category)
					categoryFound = true
					break
				}
			}
			if !categoryFound {
				return ErrCategoryNotImplemented
			}
		}
	}

	maxDate := time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)
	formattedTimeQuery := make([]time.Time, 2)
	formattedTimeQuery[0] = time.Time{}
	formattedTimeQuery[1] = maxDate
	f.Inserted = formattedTimeQuery

	if q.Has("inserted") {
		fullTimeQuery := q.Get("inserted")
		timeSplit := strings.SplitN(fullTimeQuery, ":", 2)
		for index, timeQuery := range timeSplit {
			if timeQuery != "" {
				datetime, err := time.Parse(time.DateOnly, timeQuery)
				if err != nil {
					return err
				}
				if index == 0 {
					formattedTimeQuery[index] = datetime
				} else {
					formattedTimeQuery[index] = datetime.Add(time.Hour*23 + time.Minute*59 + time.Second*59)
				}
			}
		}
		if !formattedTimeQuery[0].IsZero() && !formattedTimeQuery[1].IsZero() && formattedTimeQuery[0].After(formattedTimeQuery[1]) {
			return fmt.Errorf("start dato kan ikke være efter slut dato")
		}
		f.Inserted = formattedTimeQuery
	}

	f.Updated = formattedTimeQuery

	if q.Has("updated") {
		fullTimeQuery := q.Get("updated")
		timeSplit := strings.SplitN(fullTimeQuery, ":", 2)
		for index, timeQuery := range timeSplit {
			if timeQuery != "" {
				datetime, err := time.Parse(time.DateOnly, timeQuery)
				if err != nil {
					return err
				}
				if index == 0 {
					formattedTimeQuery[index] = datetime
				} else {
					formattedTimeQuery[index] = datetime.Add(time.Hour*23 + time.Minute*59 + time.Second*59)
				}
			}
		}
		if !formattedTimeQuery[0].IsZero() && !formattedTimeQuery[1].IsZero() && formattedTimeQuery[0].After(formattedTimeQuery[1]) {
			return fmt.Errorf("start dato kan ikke være efter slut dato")
		}
		f.Inserted = formattedTimeQuery
	}

	return nil
}

type Sender struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type Ticket struct {
	ID         int64      `json:"id" apidoc:"ignore"`
	Status     Status     `json:"status" validate:"required"`
	Categories []Category `json:"categories" validate:"required,min=1"`
	Issue      string     `json:"issue" validate:"required,max=300,min=50" description:"Description of the issue related to the device. Be as descriptive as possbile."`
	Sender     Sender     `json:"sender" validate:"required"`
	Inserted   string     `json:"inserted" apidoc:"ignore"`
	Updated    string     `json:"updated" apidoc:"ignore"`
	Logs       []Log      `json:"logs,omitempty"`
}

type Log struct {
	ID              int64  `json:"id" apidoc:"ignore"`
	TicketID        int64  `json:"ticket_id,omitempty"`
	Status          Status `json:"status"`
	Initiator       string `json:"initiator" description:"Name of the user. Sender name if triggered by ticket owner. User name if triggered by Admin user."`
	ExternalComment string `json:"external_comment"`
	InternalComment string `json:"internal_comment,omitempty"`
	Inserted        string `json:"inserted" apidoc:"ignore"`
}

type ticketStore struct {
	db *sql.DB
}

func (s *ticketStore) Create(ctx context.Context, t *Ticket) error {
	stmt := `
		INSERT INTO tickets (status, categories, issue, sender_name, sender_email, sender_address)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, inserted, updated
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		stmt,
		t.Status,
		pq.Array(t.Categories),
		t.Issue,
		t.Sender.Name,
		t.Sender.Email,
		t.Sender.Address,
	).Scan(
		&t.ID,
		&t.Inserted,
		&t.Updated,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *ticketStore) List(ctx context.Context, filters TicketFilters) ([]Ticket, int, error) {
	stmt := `
		WITH filtered_tickets AS
			(SELECT id,
					status,
					categories,
					issue,
					sender_name,
					sender_email,
					sender_address,
					inserted,
					updated
			FROM tickets
			WHERE (sender_name ILIKE '%' || $1 || '%'
					OR sender_email ILIKE '%' || $1 || '%'
					OR sender_address ILIKE '%' || $1 || '%')
				AND ($2 = 0
					  OR status = $2)
				AND ($3::int[] IS NULL
					  OR categories && $3::int[])
				AND (inserted >= $4 
					  AND inserted <= $5))
		SELECT id,
			   status,
			   categories,
			   issue,
			   sender_name,
			   sender_email,
			   sender_address,
			   inserted,
			   updated,
			   (SELECT COUNT(*) FROM filtered_tickets) AS total_count
		FROM filtered_tickets
		ORDER BY inserted DESC
		LIMIT $6
		OFFSET $7
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	tickets := []Ticket{}
	var totalCount int

	rows, err := s.db.QueryContext(
		ctx,
		stmt,
		filters.Query,
		filters.Status,
		pq.Array(filters.Categories),
		filters.Inserted[0],
		filters.Inserted[1],
		filters.Limit,
		filters.Offset,
	)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var t Ticket

		if err := rows.Scan(
			&t.ID,
			&t.Status,
			pq.Array(&t.Categories),
			&t.Issue,
			&t.Sender.Name,
			&t.Sender.Email,
			&t.Sender.Address,
			&t.Inserted,
			&t.Updated,
			&totalCount,
		); err != nil {
			return nil, 0, err
		}

		tickets = append(tickets, t)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return tickets, totalCount, nil
}

func (s *ticketStore) GetByID(ctx context.Context, ID int64) (*Ticket, error) {
	stmt := `
		SELECT id, status, categories, issue, sender_name, sender_email, sender_address, inserted, updated
		FROM tickets
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	var t Ticket

	err := s.db.QueryRowContext(
		ctx,
		stmt,
		ID,
	).Scan(
		&t.ID,
		&t.Status,
		pq.Array(&t.Categories),
		&t.Issue,
		&t.Sender.Name,
		&t.Sender.Email,
		&t.Sender.Address,
		&t.Inserted,
		&t.Updated,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrTicketNotFound
		default:
			return nil, err
		}
	}

	return &t, nil
}

func (s *ticketStore) DeleteByID(ctx context.Context, ID int64) error {
	stmt := `
		DELETE FROM tickets
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	result, err := s.db.ExecContext(
		ctx,
		stmt,
		ID,
	)
	if err != nil {
		return err
	}

	ra, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if ra == 0 {
		return ErrTicketNotFound
	}

	return nil
}

func (s *ticketStore) createLog(ctx context.Context, l *Log) error {
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

func (s *ticketStore) updateTicketStatus(ctx context.Context, ID int64, status Status) error {
	stmt := `
		UPDATE tickets
		SET status = $1, updated = $2
		WHERE id = $3
	`

	res, err := s.db.ExecContext(
		ctx,
		stmt,
		status,
		time.Now(),
		ID,
	)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if ra == 0 {
		return ErrTicketNotFound
	}

	return nil
}

func (s *ticketStore) CreateLog(ctx context.Context, l *Log) error {
	return withTx(ctx, s.db, func(tx *sql.Tx) error {
		if err := s.createLog(ctx, l); err != nil {
			return err
		}

		if err := s.updateTicketStatus(ctx, l.TicketID, l.Status); err != nil {
			return err
		}

		return nil
	})
}

func (s *ticketStore) ListInternalLogs(ctx context.Context, ID int64) ([]Log, error) {
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
		ID,
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

func (s *ticketStore) ListExternalLogs(ctx context.Context, ID int64) ([]Log, error) {
	stmt := `
		SELECT id, status, initiator, external_comment, inserted
		FROM logs
		WHERE ticket_id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	logs := []Log{}

	rows, err := s.db.QueryContext(
		ctx,
		stmt,
		ID,
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
