package model

import "github.com/google/uuid"

type User struct {
	Name     string    `db:"username" json:"name"`
	Password string    `db:"password" json:"password"`
	Email    string    `db:"email" json:"email"`
	ID       uuid.UUID `db:"id" json:"ID"`
}
