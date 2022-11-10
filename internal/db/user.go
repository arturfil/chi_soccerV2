package db

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/arturfil/gorilla_soccer/internal/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Signup - This method will create a user makind sure that the user
// we are tryin to create doesn't exist already.
func (d *Database) Signup(ctx context.Context, user user.User) (user.User, error) {
	newId := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return user, err
	}
	query := `
		insert into users(id, email, first_name, last_name, clearance, password, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8) returning id;
	`
	err = d.Client.QueryRowContext(ctx, query,
		newId,
		user.Email,
		user.FirstName,
		user.LastName,
		"member",
		hashedPassword,
		time.Now(),
		time.Now(),
	).Scan(&newId)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetUserByEmail - method that we will use in both signup and login in the controllers
// and services to see either user exists or doesn't
func (d *Database) GetUserByEmail(ctx context.Context, email string) (user.User, error) {
	query := `
		select id, email, first_name, last_name, password, created_at, updated_at from users
		where email = $1
	`
	row := d.Client.QueryRowContext(ctx, query, email)
	var user user.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

// DeleteUser - This method will be for admin purposes
func (d *Database) DeleteUser(ctx context.Context, id string) (bool, error) {
	query := `delete from users when id = $1`
	err := d.Client.QueryRowContext(ctx, query, id)
	if err != nil {
		log.Println(err)
		return false, nil
	}
	return true, nil
}

// PasswordMatches - This Method will be used to make sure that passwords match when login in
// possibly could be used in the future to check signup too
func (d *Database) PasswordMatches(plainText string, user user.User) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, err
		default:
			return false, err
		}
	}
	return true, nil
}
