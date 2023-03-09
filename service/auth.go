package service

import (
	"context"
	"time"

	"github.com/2paperstar/movie-api/config"
	"github.com/2paperstar/movie-api/database"
	"github.com/2paperstar/movie-api/model"
	"github.com/golang-jwt/jwt"
)

func RegisterWithCredential(form *model.RegisterForm) (*model.UserInfo, error) {
	return database.CreateUser(context.TODO(), form)
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
