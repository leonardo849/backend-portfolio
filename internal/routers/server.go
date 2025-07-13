package routers

import (
	_ "backend-portfolio/internal/dto"
	"backend-portfolio/internal/logger"
	"backend-portfolio/internal/middewares"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// hello godoc
// @Summary hello 
// @Description welcome message
// @Accept json
// @Produce json
// @Success 200 {object} dto.MessageResponseDTO
// @Router / [get]
func hello (ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"message": "what's up?"})
}
	

func SetupApp() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	
	logger.ZapLogger.Info("cors is ready")
	app.Use(middlewares.LogRequestsMiddleware())
	
	
	app.Get("/", hello)

	app.Get("/swagger/*", swagger.HandlerDefault)
	logger.ZapLogger.Info("swagger is ready")

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

	return app.Listen(":" + port)
}