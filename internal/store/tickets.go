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
	CREATED Status = iota + 1
	CLOSED
	RECEIVED
	EXTERNAL_REPAIR
	INTERNAL_REPAIR
	AWAITING_PARTS
	QUOTE_SENT
	QUOTE_ACCEPTED
	REJECTED
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
	CREATED:         "registreret",
	RECEIVED:        "modtaget",
	CLOSED:          "lukket",
	EXTERNAL_REPAIR: "ekstern reparation",
	INTERNAL_REPAIR: "intern reparation",
	AWAITING_PARTS:  "afventer reservedele",
	QUOTE_SENT:      "tilbud sendt",
	QUOTE_ACCEPTED:  "tilbud accepteret",
	REJECTED:        "afvist",
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
	Status     []Status    `json:"status"`
	Categories []Category  `json:"categories"`
	Inserted   []time.Time `json:"inserted"`
	Updated    []time.Time `json:"updated"`
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
	Page       int         `json:"page"`
}

func (f *TicketFilters) Parse(r *http.Request) error {
	q := r.URL.Query()

	f.Limit = 25
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

	if q.Has("page") {
		page, err := strconv.Atoi(q.Get("page"))
		if err != nil {
			return fmt.Errorf("page param kunne ikke konverteres til et nummer: %s", q.Get("page"))
		}
		f.Offset = (page - 1) * f.Limit
	}

	if q.Has("query") {
		f.Query = q.Get("query")
	}

	if q.Has("status") {
		statusQuery := strings.Split(q.Get("status"), ",")
		for _, query := range statusQuery {
			if query == "" {
				break
			}
			statusFound := false
			for status, label := range Statuses {
				if label == query {
					f.Status = append(f.Status, status)
					statusFound = true
					break
				}
			}
			if !statusFound {
				return ErrStatusNotImplemented
			}
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

type Contact struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Street  string `json:"street"`
	City    string `json:"city"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

type Ticket struct {
	ID           string     `json:"id" apidoc:"ignore"`
	Status       Status     `json:"status"`
	Categories   []Category `json:"categories"`
	Issue        string     `json:"issue" description:"Description of the issue related to the device. Be as descriptive as possbile."`
	Model        *string    `json:"model"`
	SerialNumber *string    `json:"serial_number"`
	Quote        string     `json:"quote"`
	Warranty     string     `json:"warranty"`
	Sender       Contact    `json:"sender"`
	Billing      Contact    `json:"billing"`
	Inserted     string     `json:"inserted" apidoc:"ignore"`
	Updated      string     `json:"updated" apidoc:"ignore"`
	Logs         []Log      `json:"logs,omitempty"`
}

type Log struct {
	ID              int64  `json:"id" apidoc:"ignore"`
	TicketID        string `json:"ticket_id,omitempty"`
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
	return withTx(ctx, s.db, func(tx *sql.Tx) error {
		if err := s.create(ctx, t); err != nil {
			return err
		}

		createdLog := &Log{
			TicketID:        t.ID,
			Status:          CREATED,
			Initiator:       t.Sender.Name,
			ExternalComment: "Sagen er registreret. Skancode A/S afventer modtagelse af RMA-enheden.",
			InternalComment: "",
		}

		if err := s.createLog(ctx, createdLog); err != nil {
			return err
		}

		return nil
	})
}

func (s *ticketStore) create(ctx context.Context, t *Ticket) error {
	stmt := `
		INSERT INTO tickets (
			status,
			categories,
			issue,
			model,
			serial_number,
			quote,
			warranty,
			sender_name,
			sender_email,
			sender_phone,
			sender_street,
			sender_city,
			sender_zip,
			sender_country,
			billing_name,
			billing_email,
			billing_phone,
			billing_street,
			billing_city,
			billing_zip,
			billing_country
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
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
		t.Model,
		t.SerialNumber,
		t.Quote,
		t.Warranty,
		t.Sender.Name,
		t.Sender.Email,
		t.Sender.Phone,
		t.Sender.Street,
		t.Sender.City,
		t.Sender.Zip,
		t.Sender.Country,
		t.Billing.Name,
		t.Billing.Email,
		t.Billing.Phone,
		t.Billing.Street,
		t.Billing.City,
		t.Billing.Zip,
		t.Billing.Country,
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
		WITH filtered_tickets AS (
			SELECT *
			FROM tickets
			WHERE (
				$1 = '' OR (
					sender_name ILIKE '%' || $1 || '%'
					OR sender_email ILIKE '%' || $1 || '%'
					OR sender_street ILIKE '%' || $1 || '%'
					OR sender_city ILIKE '%' || $1 || '%'
					OR sender_zip ILIKE '%' || $1 || '%'
					OR sender_country ILIKE '%' || $1 || '%'
					OR billing_name ILIKE '%' || $1 || '%'
					OR billing_email ILIKE '%' || $1 || '%'
					OR billing_street ILIKE '%' || $1 || '%'
					OR billing_city ILIKE '%' || $1 || '%'
					OR billing_zip ILIKE '%' || $1 || '%'
					OR billing_country ILIKE '%' || $1 || '%'
					OR model ILIKE '%' || $1 || '%'
					OR serial_number ILIKE '%' || $1 || '%'
					OR id ILIKE '%' || $1 || '%'
					)
				)
				AND (
					$2::integer[] IS NULL 
					OR cardinality($2::integer[]) = 0 
					OR status = ANY($2::integer[])
				)
				AND (
					$3::integer[] IS NULL 
					OR cardinality($3::integer[]) = 0 
					OR categories && $3::integer[]
				)
				AND (
					inserted BETWEEN $4 AND $5
				)
			)
		SELECT 
			id,
			status,
			categories,
			issue,
			model,
			serial_number,
			quote,
			warranty,
			sender_name,
			sender_email,
			sender_phone,
			sender_street,
			sender_city,
			sender_zip,
			sender_country,
			billing_name,
			billing_email,
			billing_phone,
			billing_street,
			billing_city,
			billing_zip,
			billing_country,
			inserted,
			updated,
			(SELECT COUNT(*) FROM filtered_tickets) AS total_count
		FROM filtered_tickets
		ORDER BY inserted DESC
		LIMIT $6 OFFSET $7
    `

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	tickets := []Ticket{}
	var totalCount int

	rows, err := s.db.QueryContext(
		ctx,
		stmt,
		filters.Query,
		pq.Array(filters.Status),
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
			&t.Model,
			&t.SerialNumber,
			&t.Quote,
			&t.Warranty,
			&t.Sender.Name,
			&t.Sender.Email,
			&t.Sender.Phone,
			&t.Sender.Street,
			&t.Sender.City,
			&t.Sender.Zip,
			&t.Sender.Country,
			&t.Billing.Name,
			&t.Billing.Email,
			&t.Billing.Phone,
			&t.Billing.Street,
			&t.Billing.City,
			&t.Billing.Zip,
			&t.Billing.Country,
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

func (s *ticketStore) GetByID(ctx context.Context, ID string) (*Ticket, error) {
	stmt := `
		SELECT
			id,
			status,
			categories,
			issue,
			model,
			serial_number,
			quote,
			warranty,
			sender_name,
			sender_email,
			sender_phone,
			sender_street,
			sender_city,
			sender_zip,
			sender_country,
			billing_name,
			billing_email,
			billing_phone,
			billing_street,
			billing_city,
			billing_zip,
			billing_country,
			inserted,
			updated
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
		&t.Model,
		&t.SerialNumber,
		&t.Quote,
		&t.Warranty,
		&t.Sender.Name,
		&t.Sender.Email,
		&t.Sender.Phone,
		&t.Sender.Street,
		&t.Sender.City,
		&t.Sender.Zip,
		&t.Sender.Country,
		&t.Billing.Name,
		&t.Billing.Email,
		&t.Billing.Phone,
		&t.Billing.Street,
		&t.Billing.City,
		&t.Billing.Zip,
		&t.Billing.Country,
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

func (s *ticketStore) DeleteByID(ctx context.Context, ID string) error {
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

func (s *ticketStore) updateTicketStatus(ctx context.Context, ID string, status Status) error {
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

func (s *ticketStore) ListInternalLogs(ctx context.Context, ID string) ([]Log, error) {
	stmt := `
		SELECT id, status, initiator, external_comment, internal_comment, inserted
		FROM logs
		WHERE ticket_id = $1
		ORDER BY id DESC
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

func (s *ticketStore) ListExternalLogs(ctx context.Context, ID string) ([]Log, error) {
	stmt := `
		SELECT id, status, initiator, external_comment, inserted
		FROM logs
		WHERE ticket_id = $1
		ORDER BY id DESC
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
