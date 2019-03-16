package user

import (
	"context"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Repository interface {
	Close() error
	GetUser(ctx context.Context, id string) (*User, error)
	PostUser(ctx context.Context, n User) error
	UpdateUser(ctx context.Context, n User) error
	DeleteUser(ctx context.Context, id string) error
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

func (r *postgresRepository) GetUser(ctx context.Context, id string) (*User, error) {
	u := &User{ID: id}
	err := r.db.Select(u)

	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *postgresRepository) PostUser(ctx context.Context, u User) error {
	user := &User{
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	err := r.db.Insert(user)
	return err
}

func (r *postgresRepository) UpdateUser(ctx context.Context, u User) error {
	user := &User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		UpdatedAt: u.UpdatedAt,
	}

	err := r.db.Update(user)
	return err
}

func (r *postgresRepository) DeleteUser(ctx context.Context, id string) error {
	user := &User{
		ID: id,
	}

	err := r.db.Delete(user)
	return err
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*User)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
