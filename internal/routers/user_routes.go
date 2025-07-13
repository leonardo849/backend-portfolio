package routers

import (
	_ "backend-portfolio/internal/dto"
	"backend-portfolio/internal/handlers"
	"backend-portfolio/internal/logger"
	"backend-portfolio/internal/repository"
	"backend-portfolio/internal/services"

	"github.com/gofiber/fiber/v2"
)

func setupUserRoutes(userGroup fiber.Router) {
	userCollection := repository.GetCollections().UsersCollection
	userRepository := repository.NewUserRepository(userCollection)
	userService := services.NewUserService(userRepository)
	userController := handlers.NewUserController(userService)

	userGroup.Post("/login", userController.LoginUser())
	logger.ZapLogger.Info("user's routes are working!")
}