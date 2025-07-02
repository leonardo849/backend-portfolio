package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func CreateJWT(id string, username string, role string,updatedAt time.Time) (string, error){
	claims := jwt.MapClaims{
		"username": username,
		"id": id,
		"updatedAt": updatedAt,
		"role": role,
		"exp": time.Now().Add(7 * (time.Hour * 24)).Unix(),
	}
	secret := os.Getenv("SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return  token.SignedString([]byte(secret))
}