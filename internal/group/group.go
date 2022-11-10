package group

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrFetchingGroup  = errors.New("failed to fetch Group by id")
	ErrNotImplemented = errors.New("not implemented")
)

// Store - defines all of the methods that our service
// needs to operatef
type GroupStore interface {
	GetGroups(context.Context) ([]Group, error)
	GetGroupById(context.Context, string) (Group, error)
	CreateGroup(context.Context, Group) (Group, error)
	UpdateGroup(context.Context, Group, string) (Group, error)
	DeleteGroup(context.Context, string) error
}

// Group - a representation of the Group
// structure
type Group struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Service - is the struct on which all our logic
// will be built on top of
type Service struct {
	Store GroupStore
}

// NewService - returns a pointer to a new service
func NewService(store GroupStore) *Service {
	return &Service{
		Store: store,
	}
}

// GetGroup - provided the id, the method returns
// the specific requested gaem
func (s *Service) GetGroups(ctx context.Context) ([]Group, error) {
	fmt.Println("return Groups")
	var groups []Group
	return groups, nil
}

func (s *Service) GetGroupById(ctx context.Context, id string) (Group, error) {
	fmt.Println("return Group")

	group, err := s.Store.GetGroupById(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Group{}, ErrFetchingGroup
	}
	return group, nil
}

func (s *Service) CreateGroup(ctx context.Context, group Group) (Group, error) {
	return Group{}, ErrNotImplemented
}

func (s *Service) UpdateGroup(ctx context.Context, id string, updatedGroup Group) (Group, error) {
	var group Group
	return group, nil
}

func (s *Service) DeleteGroup(ctx context.Context, id string) error {
	return nil
}
