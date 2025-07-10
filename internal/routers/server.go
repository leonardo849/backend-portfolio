package routers

import (
	"backend-portfolio/internal/logger"
	"backend-portfolio/internal/middewares"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
)

func SetupApp() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	logger.ZapLogger.Info("cors is ready")
	app.Use(middlewares.LogRequestsMiddleware())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{"message": "what's up?"})
	})
	userGroup := app.Group("/user")
	setupUserRoutes(userGroup)
	logger.ZapLogger.Info("app is running!")
	return  app
}

func RunServer() error {
	app := SetupApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Get("/", func (c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "welcome!"})
	})

	return app.Listen(":" + port)
}