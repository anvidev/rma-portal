package store

import (
	"context"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Hash     Password `json:"-"`
	IsActive bool     `json:"is_active"`
	Inserted string   `json:"inserted"`
}

type Password struct {
	text  *string
	bytes []byte
}

func (p *Password) Hash(str string) error {
	b, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.text = &str
	p.bytes = b
	return nil
}

func (p *Password) Compare(str string) error {
	err := bcrypt.CompareHashAndPassword(p.bytes, []byte(str))
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return fmt.Errorf("invalid credentials")
		default:
			return err
		}
	}
	return nil
}

func (p *Password) Reset() {
	p = &Password{}
}

type userStore struct {
	db *sql.DB
}

func (s *userStore) Create(ctx context.Context, u *User) error {
	defer u.Hash.Reset()

	stmt := `
		INSERT INTO users (name, email, hash)
		VALUES ($1, $2, $3)
		RETURNING id, is_active, inserted
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		stmt,
		u.Name,
		u.Email,
		u.Hash.bytes,
	).Scan(
		&u.ID,
		&u.IsActive,
		&u.Inserted,
	)

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrUserDublicateEmail
		default:
			return err
		}
	}

	return nil
}

func (s *userStore) GetByEmail(ctx context.Context, email string) (User, error) {
	stmt := `
		SELECT id, name, email, hash, is_active, inserted
		FROM users
		WHERE email = $1
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	var user User

	if err := s.db.QueryRowContext(
		ctx,
		stmt,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Hash.bytes,
		&user.IsActive,
		&user.Inserted,
	); err != nil {
		switch err {
		case sql.ErrNoRows:
			return user, ErrUserNotFound
		default:
			return user, err
		}
	}

	return user, nil
}

func (s *userStore) GetByID(ctx context.Context, ID int64) (User, error) {
	stmt := `
		SELECT id, name, email, hash, is_active, inserted
		FROM users
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	var user User

	if err := s.db.QueryRowContext(
		ctx,
		stmt,
		ID,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Hash.bytes,
		&user.IsActive,
		&user.Inserted,
	); err != nil {
		switch err {
		case sql.ErrNoRows:
			return user, ErrUserNotFound
		default:
			return user, err
		}
	}

	return user, nil
}

func (s *userStore) List(ctx context.Context) ([]User, error) {
	stmt := `
		SELECT id, name, email, is_active, inserted
		FROM users
		ORDER BY id DESC
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	var users []User

	rows, err := s.db.QueryContext(
		ctx,
		stmt,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var u User

		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.IsActive,
			&u.Inserted,
		); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
