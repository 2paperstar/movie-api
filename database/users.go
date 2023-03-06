package database

import (
	"context"
	"errors"

	"github.com/2paperstar/movie-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrDuplicatedUser = errors.New("duplicated user")

func CreateUser(ctx context.Context, payload *model.RegisterForm) (*model.UserInfo, error) {
	result, err := usersCollection.InsertOne(ctx, payload)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, ErrDuplicatedUser
		}
		return nil, err
	}
	user := usersCollection.FindOne(ctx, bson.M{"_id": result.InsertedID})
	if err := user.Err(); err != nil {
		return nil, err
	}

	userInfo := new(model.UserInfo)
	user.Decode(userInfo)
	return userInfo, nil
}
