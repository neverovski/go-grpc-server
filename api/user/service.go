package user

import (
	"context"
	"github.com/segmentio/ksuid"
	"time"
)

type Service interface {
	PostUser(ctx context.Context, username, email string) (*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s *userService) PostUser(ctx context.Context, username, email string) (*User, error) {
	n := &User{
		ID:        ksuid.New().String(),
		Username:  username,
		Email:     email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	return n, nil
}

func (s *userService) GetUser(ctx context.Context, id string) (*User, error) {
	return n, nil
}
