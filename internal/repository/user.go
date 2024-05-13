package repository

import (
	"context"

	"github.com/DarkReduX/pet-api/model"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	conn *pgxpool.Pool
}

// NewUser - new object of repository.User
func NewUser(conn *pgxpool.Pool) *User {
	return &User{conn: conn}
}

// Create creates a new user in the database.
//
// ctx: The context for the operation.
// user: A pointer to the user object to be created. The ID field of the user object will be automatically generated.
//
// Returns:
// A pointer to the created user if successful and no error occurred, nil otherwise.
// An error if the operation fails.
func (r *User) Create(ctx context.Context, user *model.User) (*model.User, error) {
	user.ID = uuid.New()

	_, err := r.conn.Exec(
		ctx,
		"INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4)",
		user.ID, user.Name, user.Email, user.Password,
	)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *User) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := new(model.User)

	err := pgxscan.Get(ctx, r.conn, user, "SELECT id, username, email, password FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *User) GetByIDs(ctx context.Context, ids []uuid.UUID) ([]*model.User, error) {
	users := make([]*model.User, len(ids))
	err := pgxscan.Select(ctx, r.conn, users, "SELECT id, username, email FROM users WHERE id = ANY($1)", ids)
	if err != nil {
		return nil, err
	}

	return users, err
}

func (r *User) Update(ctx context.Context, user *model.User) error {
	_, err := r.conn.Exec(
		ctx,
		"UPDATE users SET username=$1, email=$2, password=$3 WHERE id=$4",
		user.Name, user.Email, user.Password, user.ID,
	)
	return err
}

func (r *User) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.conn.Exec(ctx, "DELETE FROM users WHERE id = $1", id)

	return err
}
