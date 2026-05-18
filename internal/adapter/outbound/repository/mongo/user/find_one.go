package user

import (
	"context"
	"errors"

	"go-simple-template/internal/adapter/outbound/repository/mongo/model"
	"go-simple-template/internal/core/domain/entity"
	"go-simple-template/internal/core/domain/errs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *user) FindOne(ctx context.Context, user entity.User) (*entity.User, error) {
	var (
		userModel model.User
		f         bson.M
	)

	objectId, _ := primitive.ObjectIDFromHex(user.Id)

	if !objectId.IsZero() {
		f = bson.M{"_id": objectId}
	}

	if user.Email != "" {
		f = bson.M{"email": user.Email}
	}

	if f == nil {
		return nil, errs.ErrUserNotFound
	}

	result := db.collection.FindOne(ctx, f).Decode(&userModel)
	if result != nil {
		if errors.Is(result, mongo.ErrNoDocuments) {
			return nil, errs.ErrUserNotFound
		}

		return nil, result
	}

	return userModel.ToEntity(), nil
}
