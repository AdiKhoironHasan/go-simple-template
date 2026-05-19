package user

import (
	"context"
	"errors"

	"go-simple-template/internal/adapter/outbound/repository/mongo/model"
	"go-simple-template/internal/core/domain/entity"

	errpkg "go-simple-template/internal/pkg/errs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *user) FindOne(ctx context.Context, user entity.User) (*entity.User, error) {
	var (
		userModel model.User
		f         bson.M
	)

	if user.Id != "" {
		objectId, err := primitive.ObjectIDFromHex(user.Id)
		if err != nil {
			return nil, errpkg.NewNotFound("invalid user id")
		}
		f = bson.M{"_id": objectId}
	}

	if user.Email != "" {
		f = bson.M{"email": user.Email}
	}

	if f == nil {
		return nil, errpkg.NewNotFound("user not found")
	}

	result := db.collection.FindOne(ctx, f).Decode(&userModel)
	if result != nil {
		if errors.Is(result, mongo.ErrNoDocuments) {
			return nil, errpkg.NewNotFound("user not found")
		}

		return nil, result
	}

	return userModel.ToEntity(), nil
}
