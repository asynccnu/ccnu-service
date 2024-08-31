package data

import (
	"context"
	"github.com/asynccnu/ccnu-service/internal/biz"
)

type UserRepo struct {
	data *Data
}

func NewUserRepo(data *Data) *UserRepo {
	return &UserRepo{
		data: data,
	}
}

func (r *UserRepo) Save(ctx context.Context, user *biz.User) error {
	db := r.data.DB.Table(biz.UserTableName).WithContext(ctx)
	return db.FirstOrCreate(&user).Error
}

func (r *UserRepo) GetByUserID(ctx context.Context, userID string) (*biz.User, error) {
	var user biz.User
	db := r.data.DB.Table(biz.UserTableName).WithContext(ctx)
	err := db.Where("userid = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
