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
var ErrInvalidToken = errors.New("invalid token")

type authTokenPayload struct {
	UID                string `json:"uid"`
	jwt.StandardClaims `json:",inline"`
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
	accessTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS512, authTokenPayload{
		UID: user.UID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}).SignedString(config.JwtSecret)
	if err != nil {
		return nil, err
	}
	refreshTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS512, authTokenPayload{
		UID: user.UID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}).SignedString(config.JwtSecret)
	if err != nil {
		return nil, err
	}
	return &model.AuthResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func VerifyJwt(tokenString string) (*authTokenPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return config.JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	authPayload := new(authTokenPayload)
	authPayload.UID = payload["uid"].(string)
	authPayload.ExpiresAt = int64(payload["exp"].(float64))

	return authPayload, nil
}
