package integration_test

import (
	"backend-portfolio/internal/logger"
	"os"
	"testing"

	"go.uber.org/zap"
)

func TestLoginUser(t *testing.T) {
	username := "jhon_doe"
	password := os.Getenv("PASSWORD_ADMIN")
	if password == "" {
		logger.ZapLogger.Error(
			"there isn't password admin",
			zap.String("function", "login user test"),
		)
		os.Exit(1)
	}
	e := newExpect(t)
	e.POST("/user/login"). 
	WithJSON(map[string]string{
		"username": username,
		"password": password,
	}). 
	Expect(). 
	Status(200)
}