package service

import (
	"context"
	"errors"

	"github.com/2paperstar/movie-api/database"
	"github.com/2paperstar/movie-api/model"
)

var ErrDuplicatedUser = errors.New("duplicated user")

func RegisterWithCredential(form *model.RegisterForm) (*model.UserInfo, error) {
	if _, _, err := database.GetUserByLoginID(context.TODO(), form.ID); err == nil {
		return nil, ErrDuplicatedUser
	}
	return database.CreateUser(context.TODO(), form)
}

func UpdateUser(uid string, form *model.UserInfoForm) (*model.UserInfo, error) {
	return database.UpdateUser(context.TODO(), uid, form)
}
