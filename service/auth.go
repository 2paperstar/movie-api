package service

import (
	"context"

	"github.com/2paperstar/movie-api/database"
	"github.com/2paperstar/movie-api/model"
)

func RegisterWithCredential(form *model.RegisterForm) (*model.UserInfo, error) {
	return database.CreateUser(context.TODO(), form)
}
