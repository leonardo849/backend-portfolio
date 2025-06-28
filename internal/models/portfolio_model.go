package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PortfolioModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Topics []string `json:"topic" bson:"topic"`
	Projects []ProjectModel `bson:"projects" json:"projects"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}