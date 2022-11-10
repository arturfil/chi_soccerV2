package field

import (
	"context"
	"fmt"
	"time"
)

type FieldStore interface {
	GetFields(context.Context) ([]Field, error)
	CreateField(context.Context, Field) (Field, error)
}

type Field struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Service struct {
	Store FieldStore
}

func NewService(store FieldStore) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetFields(ctx context.Context) ([]Field, error) {
	fields, err := s.Store.GetFields(ctx)
	if err != nil {
		fmt.Println(err)
		return []Field{}, err
	}
	return fields, nil
}

func (s *Service) CreateField(ctx context.Context, field Field) (Field, error) {
	field, err := s.Store.CreateField(ctx, field)
	if err != nil {
		return Field{}, nil
	}
	return field, nil
}
