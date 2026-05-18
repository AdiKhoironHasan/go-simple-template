package user

import (
	"context"

	"go-simple-template/internal/adapter/outbound/repository/mongo/model"
	"go-simple-template/internal/core/domain/entity"
	"go-simple-template/internal/pkg/errs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *user) Update(ctx context.Context, user entity.User) (*entity.User, error) {
	var (
		response *entity.User
	)

	objectId, _ := primitive.ObjectIDFromHex(user.Id)
	f := bson.M{
		"_id": objectId,
	}

	userModel := model.ToUserModel(&user)
	userModel.BeforeUpdate()

	update := bson.M{
		"$set": userModel,
	}

	result, err := db.collection.UpdateOne(ctx, f, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		return nil, errs.NewNotFound("user not found")
	}

	response = userModel.ToEntity()
	response.Id = objectId.Hex()
	return response, nil
}
