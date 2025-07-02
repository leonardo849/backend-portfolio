package dto

import (
	"backend-portfolio/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginUserDTO struct {
	Username string `bson:"username" json:"username" validate:"required,min=8,max=25"`
	Password string `json:"password" validate:"required,min=8,max=25"`
}

type FindUserDTO struct {
	ID           primitive.ObjectID        `bson:"_id,omitempty" json:"id"`
	PhotoURL     string                    `bson:"photo_url" json:"photo_url"`
	Username     string                    `json:"username"`
	Password *string `json:"password"`
	Slogan       string                    `bson:"slogan" json:"slogan"`
	Description  string                    `bson:"description" json:"description"`
	Skills       []string                  `json:"skills" bson:"skills"`
	Formations   []models.FormationModel   `json:"formations" bson:"formations"`
	SocialMedias []models.SocialMediaModel `json:"social_medias" bson:"social_medias"`
	Role         string                    `bson:"role" json:"role"`
	CreatedAt    time.Time                 `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time                 `bson:"updatedAt" json:"updatedAt"`
}
