package news

import (
	"context"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Repository interface {
	Close() error
	GetNews(ctx context.Context, id string) (*News, error)
	GetNewsForUser(ctx context.Context, userID string) ([]News, error)
	PostNews(ctx context.Context, n News) error
	PutNews(ctx context.Context, n News) error
	DeleteNews(ctx context.Context, id string) error
}

type postgresRepository struct {
	db *pg.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db := pg.Connect(&pg.Options{
		User: "postgres",
	})

	defer db.Close()

	err := createSchema(db)
	if err != nil {
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() error {
	return r.db.Close()
}

func (r *postgresRepository) GetNews(ctx context.Context, id string) (*News, error) {
	news := &News{
		ID: id,
	}

	err := r.db.Select(news)
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (r *postgresRepository) GetNewsForUser(ctx context.Context, userID string) ([]News, error) {
	var news []News

	err := r.db.Model(&news).Where("news.user_id = ?", userID).Select()

	if err != nil {
		return nil, err
	}
	return news, nil
}

func (r *postgresRepository) PostNews(ctx context.Context, n News) error {
	news := &News{
		Title:       n.Title,
		Description: n.Description,
		H1:          n.H1,
		Text:        n.Text,
		Published:   n.Published,
		CreatedAt:   n.CreatedAt,
		UpdatedAt:   n.UpdatedAt,
	}

	err := r.db.Insert(news)
	return err
}

func (r *postgresRepository) PutNews(ctx context.Context, n News) error {
	news := &News{
		ID:          n.ID,
		Title:       n.Title,
		Description: n.Description,
		H1:          n.H1,
		Text:        n.Text,
		Published:   n.Published,
		UpdatedAt:   n.UpdatedAt,
	}

	err := r.db.Update(news)
	return err
}

func (r *postgresRepository) DeleteNews(ctx context.Context, id string) error {
	news := &News{
		ID: id,
	}

	err := r.db.Delete(news)
	return err
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*News)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
