package model

import (
	"go-simple-template/internal/core/domain/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Base     `bson:"inline"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func (User) CollectionName() string {
	return "user"
}

func ToUserModel(m *entity.User) *User {
	id, _ := primitive.ObjectIDFromHex(m.Id)

	return &User{
		Base: Base{
			Id:        id,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
		Name:     m.Name,
		Email:    m.Email,
		Password: m.Password,
	}
}

func (m *User) ToEntity() *entity.User {
	return &entity.User{
		Base: entity.Base{
			Id:        m.Id.Hex(),
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
		Name:     m.Name,
		Email:    m.Email,
		Password: m.Password,
	}
}
