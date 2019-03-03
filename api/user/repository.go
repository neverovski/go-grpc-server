package user

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository interface {
	Close() error
	Ping() error
	GetUser(ctx context.Context, id string) (*User, error)
	PostUser(ctx context.Context, n User) error
	UpdateUser(ctx context.Context, n User) error
	DeleteUser(ctx context.Context, id string) error
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

func (r *postgresRepository) Close() error {
	return r.db.Close()
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}

func (r *postgresRepository) GetUser(ctx context.Context, id string) (*User, error) {
	sqlStatement := "SELECT * FROM users WHERE id = $1"

	row := r.db.QueryRowContext(ctx, sqlStatement, id)
	u := &User{}
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.UpdatedAt, &u.CreatedAt); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *postgresRepository) PostUser(ctx context.Context, u User) error {
	sqlStatement := "INSERT INTO users(id, username, email, created_at, updated_at) VALUES($1, $2, $3, $4, $5)"

	_, err := r.db.ExecContext(
		ctx,
		sqlStatement,
		u.ID,
		u.Username,
		u.Email,
		u.CreatedAt,
		u.UpdatedAt,
	)
	return err
}

func (r *postgresRepository) UpdateUser(ctx context.Context, u User) error {
	sqlStatement := "UPDATE users SET username = $2, email = $3, updated_at = $4 WHERE id = $1"

	_, err := r.db.Exec(sqlStatement, u.ID, u.Username, u.Email, u.UpdatedAt)
	return err
}

func (r *postgresRepository) DeleteUser(ctx context.Context, id string) error {
	sqlStatement := "DELETE FROM users WHERE id = ?"

	_, err := r.db.Exec(
		sqlStatement,
		id,
	)
	return err
}
