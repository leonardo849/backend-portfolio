package main

import (
	"backend-portfolio/config"
	"backend-portfolio/internal/logger"
	"backend-portfolio/internal/repository"
	"backend-portfolio/internal/routers"
	"backend-portfolio/internal/validate"
	"log"
	"os"

	"go.uber.org/zap"
)

// @title Backend Portfolio API
// @version 1.0
// @description api for a portfolio project
// @host localhost:port
// @BasePath /
func main() {
	if err := config.SetupEnvVar(); err != nil {
		log.Fatal(err.Error())
	}
	if err := logger.StartLogger(); err != nil {
		log.Fatal(err.Error())
	}
	if _,err := repository.ConnectToDatabase(); err != nil {
		os.Exit(1)
	}
	validate.StartValidator()
	if err := routers.RunServer(); err != nil {
		logger.ZapLogger.Error("error in run server", 
		zap.Error(err),
		zap.String("function", "routers.RunServer()"),
		)
	}
}