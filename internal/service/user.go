package service

import (
	"context"
	"petProject/internal/repository"
	"petProject/model"

	"github.com/google/uuid"
)

type User struct {
	rep *repository.User
}

func NewUser(rep *repository.User) *User {
	return &User{rep: rep}
}

func (u *User) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return u.rep.GetByEmail(ctx, email)
}

func (u *User) Delete(ctx context.Context, id uuid.UUID) error {
	return u.rep.Delete(ctx, id)
}

func (u *User) Update(ctx context.Context, user *model.User) error {
	return u.rep.Update(ctx, user)
}
