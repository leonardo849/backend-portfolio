package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SocialMediaModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Name string `json:"name" bson:"name"`
	Url string `json:"url" bson:"url"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}