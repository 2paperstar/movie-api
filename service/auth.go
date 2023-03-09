package service

import (
	"context"
	"errors"
	"time"

	"github.com/2paperstar/movie-api/config"
	"github.com/2paperstar/movie-api/database"
	"github.com/2paperstar/movie-api/model"
	"github.com/golang-jwt/jwt"
)

var ErrLoginFailed = errors.New("login failed")

func RegisterWithCredential(form *model.RegisterForm) (*model.UserInfo, error) {
	return database.CreateUser(context.TODO(), form)
}

func AuthorizeWithCredential(form *model.Credential) (*model.UserInfo, error) {
	user, credential, err := database.GetUserByLoginID(context.TODO(), form.ID)
	if err != nil {
		return nil, ErrLoginFailed
	}

	if credential.Password != form.Password {
		return nil, ErrLoginFailed
	}
	return user, nil
}

func GenerateJwt(user *model.UserInfo) (*model.AuthResponse, error) {
	accessTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		Id:        user.UID,
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	}).SignedString(config.JwtSecret)
	if err != nil {
		return nil, err
	}
	refreshTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		Id:        user.UID,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
	}).SignedString(config.JwtSecret)
	if err != nil {
		return nil, err
	}
	return &model.AuthResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
