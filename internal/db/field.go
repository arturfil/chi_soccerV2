package db

import (
	"context"
	"log"
	"time"

	"github.com/arturfil/gorilla_soccer/internal/field"
	"github.com/google/uuid"
)

type FieldRow struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func convertFieldRowToField(f FieldRow) field.Field {
	return field.Field{
		ID:        f.ID,
		Name:      f.Name,
		Address:   f.Address,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}

func (d *Database) GetFields(ctx context.Context) ([]field.Field, error) {
	query := `select * from fields`
	rows, err := d.Client.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var fields []field.Field
	for rows.Next() {
		var field field.Field
		err := rows.Scan(
			&field.ID,
			&field.Name,
			&field.Address,
			&field.CreatedAt,
			&field.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		fields = append(fields, field)
	}
	return fields, nil
}

func (d *Database) CreateField(ctx context.Context, field field.Field) (field.Field, error) {
	newId := uuid.New()
	query := `
		insert into fields (id, name, address, created_at, updated_at)
		values ($1, $2, $3, $4, $5) returning id;
	`

	err := d.Client.QueryRowContext(ctx, query,
		newId,
		field.Name,
		field.Address,
		time.Now(),
		time.Now(),
	).Scan(&newId)
	if err != nil {
		log.Print(err)
		return field, nil
	}

	return field, nil
}
