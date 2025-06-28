package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PortfolioID primitive.ObjectID `bson:"portfolio_id,omitempty" json:"portfolio_id"`
	Title string `json:"title" bson:"title"`
	Url string `json:"url" bson:"url"`
	Description string `json:"description" bson:"description"`
	UsedTools []string `json:"used_tools" bson:"used_tools"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}