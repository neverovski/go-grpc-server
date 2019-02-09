package user

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository interface {
	Close()
	Ping() error
	PutUser(ctx context.Context, n User) error
	GetUser(ctx context.Context, id string) (*User, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	r.db.Close()
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}

func (r *postgresRepository) PutUser(ctx context.Context, u User) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO news(id, username, email, created_at, updated_at) VALUES($1, $2, $3, $4, $5)",
		u.ID,
		u.Username,
		u.Email,
		u.CreatedAt,
		u.UpdatedAt,
	)
	return err
}

func (r *postgresRepository) GetUser(ctx context.Context, id string) (*User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = $1", id)
	u := &User{}
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.UpdatedAt, &u.CreatedAt); err != nil {
		return nil, err
	}
	return u, nil
}
