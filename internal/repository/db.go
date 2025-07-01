package repository

import (
	"backend-portfolio/internal/helper"
	"backend-portfolio/internal/logger"
	"backend-portfolio/internal/models"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	
	err = createJhonDoeAccount()
	if err != nil {
		return nil, err
	}
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

func createJhonDoeAccount() error {
	ctx, cancel := helper.CreateCtx()
	defer cancel()

	collections := GetCollections()

	password := os.Getenv("PASSWORD_ADMIN")
	if password == "" {
		err := fmt.Errorf("password_admin doesn't exist")
		logger.ZapLogger.Error(
			"error in get password_admin",
			zap.Error(err),
			zap.String("function", "createJhonDoeAccount"),
		)
		return err
	}

	const username = "jhon_doe"
	if err := collections.UsersCollection.FindOne(ctx, bson.M{"username": username}).Err(); err == nil {
		logger.ZapLogger.Info(
			"jhon doe already exists",
			zap.String("function", "createJhonDoeAccount"),
		)
		return nil
	}

	hash, err := helper.StringToHash(password)
	if err != nil {
		logger.ZapLogger.Error(
			"error in create hash",
			zap.Error(err),
			zap.String("function", "createJhonDoeAccount"),
		)
		return err
	}

	jhonDoe := models.UserModel{
		PhotoURL:   "https://static.wikia.nocookie.net/roblox/images/a/a8/JohnDoeROBLOXAvatar2011.png/revision/latest/scale-to-width/360?cb=20180714112000",
		Password:   hash,
		Username:   username,
		Slogan:     "I'm the first account XD",
		Description:"What are you doing here bro?",
		Skills:     []string{"absolute", "nothing"},
		Role:       "ADMIN",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	result, err := collections.UsersCollection.InsertOne(ctx, jhonDoe)
	if err != nil {
		logger.ZapLogger.Error(
			"error inserting jhon doe",
			zap.Error(err),
			zap.String("function", "createJhonDoeAccount"),
		)
		return err
	}

	idJhonDoe, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		logger.ZapLogger.Error(
			"error converting insertedId to ObjectID",
			zap.String("function", "createJhonDoeAccount"),
		)
		return fmt.Errorf("error convert insertedId to ObjectID")
	}

	dateStr := "2022-12-18"
	layout := "2006-01-02" 
	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		logger.ZapLogger.Error(
			"error parsing date",
			zap.Error(err),
			zap.String("function", "createJhonDoeAccount"),
		)
		return err
	}

	formation := models.FormationModel{
		UserID:    idJhonDoe,
		Type:      "Bachelor's degree in being a Messi fan",
		Date:      parsedDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = collections.FormationsCollection.InsertOne(ctx, formation)
	if err != nil {
		logger.ZapLogger.Error(
			"error inserting formation",
			zap.Error(err),
			zap.String("function", "createJhonDoeAccount"),
		)
		return err
	}

	logger.ZapLogger.Info("jhon doe account and formation created")
	return nil
}
