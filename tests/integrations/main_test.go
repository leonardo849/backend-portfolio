package integration_test

import (
	"backend-portfolio/config"
	"backend-portfolio/internal/helper"
	"backend-portfolio/internal/logger"
	"backend-portfolio/internal/repository"
	"backend-portfolio/internal/routers"
	"backend-portfolio/internal/validate"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var app *fiber.App


type fiberRoundTripper struct {
	app *fiber.App
}





func (rt fiberRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return rt.app.Test(req)
}

var DB *mongo.Database

func TestMain(m *testing.M) {
	err := config.SetupEnvVar()
	if err != nil {
		log.Panic(err.Error())
	}
	if err = logger.StartLogger(); err != nil {
		log.Panic(err.Error())
	}
	c, err := repository.ConnectToDatabase()
	defer cleanDatabase(DB)
	if err != nil {
		log.Panic(err.Error())
	}
	validate.StartValidator()
	
	DB = repository.DB
	app = routers.SetupApp()
	code := m.Run()
	ctx, cancel := helper.CreateCtx()
	defer cancel()
	
	c.Disconnect(ctx)
	os.Exit(code)
}

func newExpect(t *testing.T) *httpexpect.Expect {
	client := &http.Client{
		Transport: fiberRoundTripper{app: app},
	}
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost",
		Client:   client,
		Reporter: httpexpect.NewRequireReporter(t),
	})
}

func cleanDatabase(db *mongo.Database) {
	ctx, cancel := helper.CreateCtx()
	defer cancel()
	collections, err := db.ListCollectionNames(ctx, struct{}{}) 
	if err != nil {
		logger.ZapLogger.Error(
			"error in list collections names",
			zap.Error(err),
			zap.String("function", "clean database"),
		)
		os.Exit(1)
	}

	for _, collectionName := range collections {
		collection := db.Collection(collectionName)
		_, err := collection.DeleteMany(ctx, bson.M{})
		if err != nil {
			logger.ZapLogger.Error(
				"error in clean database",
				zap.Error(err),
				zap.String("function", "clean database"),
			)
			os.Exit(1)
		}
		message := "All elements of" + " " + collectionName +" " + "were deleted"
		logger.ZapLogger.Info(message)
	}
}