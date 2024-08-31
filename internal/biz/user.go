package biz

import (
	"context"
	"gorm.io/gorm"
	"time"
)

const UserTableName = "users"

type User struct {
	UserID    string `gorm:"primaryKey;column:userid" json:"userid"`
	Username  string `gorm:"column:username" json:"username"`
	Password  string `gorm:"column:password" json:"password"`
	CreatedAt time.Time
}

type UserRepo interface {
	Save(ctx context.Context, user *User) error
	GetByUserID(ctx context.Context, userID string) (*User, error)
}

func (u *User) TableName() string {
	return UserTableName
}
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) SaveUser(ctx context.Context, user *User) error {
	return uc.repo.Save(ctx, user)
}

func (uc *UserUsecase) GetUserByIDFromDB(ctx context.Context, userID string) (*User, error) {
	return uc.repo.GetByUserID(ctx, userID)
}
