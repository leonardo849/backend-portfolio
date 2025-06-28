package repository

import (
	"backend-portfolio/internal/logger"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var c *mongo.Client
var DB *mongo.Database


type CollectionsReturn struct {
	UsersCollection *mongo.Collection
	PortfoliosCollection *mongo.Collection
	FormationsCollection *mongo.Collection
	ProjectsCollection *mongo.Collection
	SocialMediaCollection *mongo.Collection
	Client *mongo.Client
}
var collectionsArrayNames = [5]string{"users", "portfolios", "formations", "projects", "social_media"}

func ConnectToDatabase() (*mongo.Client, error) {

	UriDatabase := os.Getenv("URI_DB")

	clientOptions := options.Client().ApplyURI(UriDatabase)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		logger.ZapLogger.Error("error in mongo connect"  ,zap.Error(err), zap.String("function", "connect to database"))
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.ZapLogger.Error("error in mongo ping"  ,zap.Error(err), zap.String("function", "connect to database"))
		return nil, err
	}

	logger.ZapLogger.Info("database connection is ready!")

	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		logger.ZapLogger.Error("error in get database's name"  ,zap.Error(err), zap.String("function", "connect to database"))
		return nil, fmt.Errorf("there isn't database's name")
	}
	database := client.Database(databaseName)

	
	
	for _, name := range collectionsArrayNames {
		if err := ensureCollectionExists(ctx, name, database); err != nil {
			log.Panic(err.Error())
			return nil, err
		}
	}

	DB = database
	c = client
	
	return client, nil
}


func ensureCollectionExists(ctx context.Context, nameCollection string, db *mongo.Database) error {
	collections, err := db.ListCollectionNames(ctx, bson.M{"name": nameCollection})
	if err != nil {
		return err
	}

	for _, name := range collections {
		if strings.EqualFold(name, nameCollection) {
			fmt.Println("Collection:", name, "already exists")
			return nil
		}
	}

	if err := db.CreateCollection(ctx, nameCollection); err != nil {
		return err
	}
	fmt.Println("Collection:", nameCollection, "created successfully")
	return nil
}

func GetCollections() CollectionsReturn {
	collections := CollectionsReturn{
		UsersCollection: DB.Collection(collectionsArrayNames[0]),
		PortfoliosCollection: DB.Collection(collectionsArrayNames[1]),
		Client: c,
		FormationsCollection: DB.Collection(collectionsArrayNames[2]),
		ProjectsCollection: DB.Collection(collectionsArrayNames[3]),
		SocialMediaCollection: DB.Collection(collectionsArrayNames[4]),
	}
	return collections
}