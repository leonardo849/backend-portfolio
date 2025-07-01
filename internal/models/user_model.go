package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PhotoURL     string             `bson:"photo_url" json:"photo_url"`
	Username     string             `bson:"username" json:"username"`
	Password     string             `bson:"password" json:"password"`
	Slogan       string             `bson:"slogan" json:"slogan"`
	Description  string             `bson:"description" json:"description"`
	Skills       []string           `json:"skills" bson:"skills"`
	Formations []FormationModel `json:"formations" bson:"formations"`
	SocialMedias []SocialMediaModel `json:"social_medias" bson:"social_medias"`
	Role         string             `bson:"role" json:"role"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
