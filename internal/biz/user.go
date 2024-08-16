package biz

import (
	"context"
)

type User struct {
	UserID   string
	Username string
	Password string
}

type UserRepo interface {
	Save(ctx context.Context, user *User) error
	GetByUserID(ctx context.Context, userID string) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) Save(ctx context.Context, user *User) error {
	return uc.repo.Save(ctx, user)
}

func (uc *UserUsecase) GetByUserID(ctx context.Context, userID string) (*User, error) {
	return uc.repo.GetByUserID(ctx, userID)
}
