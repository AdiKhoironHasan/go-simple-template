package user

import (
	"context"

	"go-simple-template/internal/adapter/outbound/repository/mongo/model"
	"go-simple-template/internal/core/domain/entity"
	"go-simple-template/internal/pkg/errs"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *user) Insert(ctx context.Context, user entity.User) (*entity.User, error) {
	var (
		response *entity.User
	)

	userModel := model.ToUserModel(&user)
	userModel.BeforeInsert()

	result, err := db.collection.InsertOne(ctx, userModel)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, errs.NewConflict("user already exists")
		}
		return nil, errs.NewInternal(err, "failed to insert user")
	}

	id := result.InsertedID.(primitive.ObjectID)
	if id.IsZero() {
		return nil, errs.NewInternal(nil, "failed to insert user")
	}

	response = userModel.ToEntity()
	response.Id = id.Hex()

	return response, nil
}
