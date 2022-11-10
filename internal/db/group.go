package db

import (
	"context"
	"time"

	"github.com/arturfil/gorilla_soccer/internal/group"
)

type GroupRow struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func convertGroupRowToGroup(g GroupRow) group.Group {
	return group.Group{
		ID:        g.ID,
		Name:      g.Name,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}

func (d *Database) GetGroups(ctx context.Context) ([]group.Group, error) {
	var groups []group.Group
	return groups, nil
}

func (d *Database) GetGroupById(ctx context.Context, uuid string) (group.Group, error) {
	var group group.Group
	return group, nil
}

func (d *Database) CreateGroup(ctx context.Context, group group.Group) (group.Group, error) {
	// var group group.Group
	return group, nil
}

func (d *Database) UpdateGroup(ctx context.Context, group group.Group, uuid string) (group.Group, error) {
	return group, nil
}

func (d *Database) DeleteGroup(ctx context.Context, uuid string) error {
	return nil
}
