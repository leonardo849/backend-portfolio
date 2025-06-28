package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FormationModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type string `bson:"type" json:"type"`
	Date time.Time `bson:"date" json:"date"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}