package user

import (
	"context"
	"fmt"
)

type UserStore interface {
	GetUsers(context.Context) ([]User, error)
} // interface of which methods this will have

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Service struct {
	Store UserStore
}

func NewService(store UserStore) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetUsers(ctx context.Context) ([]User, error) {
	users, err := s.Store.GetUsers(ctx)
	if err != nil {
		fmt.Println(err)
		return []User{}, err
	}
	return users, nil
}
