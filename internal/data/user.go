package data

import (
	"ccnu-service/internal/biz"
	"context"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (r *userRepo) Save(ctx context.Context, user *biz.User) error {
	return r.data.DB.WithContext(ctx).Create(&user).Error
}

func (r *userRepo) GetByUserID(ctx context.Context, userID string) (*biz.User, error) {
	var user biz.User
	err := r.data.DB.WithContext(ctx).Where("userid = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
