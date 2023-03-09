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

	return GetUserByUID(ctx, result.InsertedID)
}

func GetUserByUID(ctx context.Context, uid any) (*model.UserInfo, error) {
	user := usersCollection.FindOne(ctx, bson.M{"_id": uid})
	if err := user.Err(); err != nil {
		return nil, err
	}

	userInfo := new(model.UserInfo)
	user.Decode(userInfo)
	return userInfo, nil
}

func GetUserByLoginID(ctx context.Context, id string) (*model.UserInfo, *model.Credential, error) {
	user := usersCollection.FindOne(ctx, bson.M{"id": id})
	if err := user.Err(); err != nil {
		return nil, nil, err
	}

	userInfo := new(model.UserInfo)
	credential := new(model.Credential)
	user.Decode(userInfo)
	user.Decode(credential)
	return userInfo, credential, nil
}
