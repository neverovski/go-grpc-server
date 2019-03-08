package news

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository interface {
	Close() error
	Ping() error
	GetNews(ctx context.Context, id string) (*News, error)
	GetNewsForUser(ctx context.Context, userID string) ([]News, error)
	PostNews(ctx context.Context, n News) error
	PutNews(ctx context.Context, n News) error
	DeleteNews(ctx context.Context, id string) error
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

func (r *postgresRepository) GetNews(ctx context.Context, id string) (*News, error) {
	sqlStatement := "SELECT * FROM news WHERE id = $1"

	row := r.db.QueryRowContext(ctx, sqlStatement, id)
	n := &News{}
	if err := row.Scan(&n.ID, &n.Title, &n.Description, &n.H1, &n.Text, &n.UserID, &n.Published, &n.CreatedAt, &n.UpdatedAt); err != nil {
		return nil, err
	}
	return n, nil
}

func (r *postgresRepository) GetNewsForUser(ctx context.Context, userID string) ([]News, error) {
	sqlStatement := "SELECT n.id, n.title, n.description, n.h1, n.text, n.user_id, n.published, n.created_at, n.updated_at FROM news n WHERE n.user_id = $1 ORDER BY n.id"

	rows, err := r.db.QueryContext(
		ctx,
		sqlStatement,
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

func (r *postgresRepository) PostNews(ctx context.Context, n News) error {
	sqlStatement := "INSERT INTO news(id, title, description, h1, text, user_id, published, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	// Insert news
	_, err := r.db.ExecContext(
		ctx,
		sqlStatement,
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

func (r *postgresRepository) PutNews(ctx context.Context, n News) error {
	sqlStatement := "UPDATE news SET title = $2, description = $3, h1 = $3, text = $4, published = $5, updated_at = $6 WHERE id = $1"

	_, err := r.db.Exec(sqlStatement, n.ID, n.Title, n.Description, n.H1, n.Text, n.Text, n.Published, n.UpdatedAt)
	return err
}

func (r *postgresRepository) DeleteNews(ctx context.Context, id string) error {
	sqlStatement := "DELETE FROM news WHERE id = ?"

	_, err := r.db.Exec(
		sqlStatement,
		id,
	)
	return err
}
