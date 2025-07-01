package helper

import "golang.org/x/crypto/bcrypt"

func StringToHash(password string) (string, error) {
	passwordBytes := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}