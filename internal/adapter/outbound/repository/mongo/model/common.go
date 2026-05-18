package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (m *Base) BeforeInsert() {
	if m.Id.IsZero() {
		m.Id = primitive.NewObjectID()
	}

	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now().UTC()
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = time.Now().UTC()
	}
}

func (m *Base) BeforeUpdate() {
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = time.Now().UTC()
	}
}
