package news

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository interface {
	Close()
	PutNews(ctx context.Context, n News) error
	GetNews(ctx context.Context, id string) (*News, error)
	GetNewsForUser(ctx context.Context, userID string) ([]News, error)
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

func (r *postgresRepository) PutNews(ctx context.Context, n News) error {
	// Insert news
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO news(id, title, description, h1, text, user_id, published, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		n.ID,
		n.Title,
		n.Description,
		n.H1,
		n.Text,
		n.UserID,
		n.Published,
		n.CreatedAt,
		n.UpdatedAt,
	)
	return err
}

func (r *postgresRepository) GetNews(ctx context.Context, id string) (*News, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM news WHERE id = $1", id)
	n := &News{}
	if err := row.Scan(&n.ID, &n.Title, &n.Description, &n.H1, &n.Text, &n.UserID, &n.Published, &n.CreatedAt, &n.UpdatedAt); err != nil {
		return nil, err
	}
	return n, nil
}

func (r *postgresRepository) GetNewsForUser(ctx context.Context, userID string) ([]News, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT
      n.id,
      n.title,
      n.description,
      n.h1,
      n.text,
      n.user_id,
      n.published,
      n.created_at,
      n.updated_at,
    FROM news n WHERE n.user_id = $1
    ORDER BY n.id`,
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	news := []News{}
	for rows.Next() {
		n := &News{}
		if err = rows.Scan(&n.ID, &n.Title, &n.Description, &n.H1, &n.Text, &n.UserID, &n.Published, &n.CreatedAt, &n.UpdatedAt); err == nil {
			news = append(news, *n)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return news, nil
}
