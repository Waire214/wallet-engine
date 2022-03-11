package mongodb

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"wallet/internal/core/helper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo(dbType, dbUsername, dbPassword, dbHost, dbPort, authdb, dbname, server string) (*mongo.Database, error) {
	helper.LogEvent("INFO", "Establishing mongoDB connection with given credentials...")
	var conn *mongo.Database
	var mongoCredentials, authSource string

	if server == "dev" {
		MONGODB_URI := fmt.Sprint(dbType, "://", mongoCredentials, dbHost, ":", dbPort, "/?", authSource, "directConnection=true&serverSelectionTimeoutMS=2000")
		clientOptions := options.Client().ApplyURI(MONGODB_URI)

		helper.LogEvent("INFO", "Connecting to MongoDB...")
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			helper.LogEvent("ERROR", helper.ErrorMessage(helper.MONGO_DB_ERROR, "TEST"+err.Error()))
			return &mongo.Database{}, err
		}

		// Check the connection
		helper.LogEvent("INFO", "Confirming MongoDB Connection...")
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			helper.LogEvent("ERROR", helper.ErrorMessage(helper.MONGO_DB_ERROR, err.Error()))
			return &mongo.Database{}, err
		}

		//helper.LogEvent("Info", "Connected to MongoDB!")
		helper.LogEvent("INFO", "Establishing Database collections and indexes...")

		conn = client.Database(dbname)
		return conn, nil
	} else if server == "prod" {
		serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		// "mongodb+srv://oluwatosin3:Kewedeola9363@waire.nxjve.mongodb.net/transmed?retryWrites=true&w=majority"
		clientOptions := options.Client().
			// ApplyURI(os.Getenv("MONGODB_URI")).SetServerAPIOptions(serverAPIOptions)
			ApplyURI(helper.Config.MONGODB_URI).SetServerAPIOptions(serverAPIOptions)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			helper.LogEvent("ERROR", helper.ErrorMessage(helper.MONGO_DB_ERROR, "TEST"+err.Error()))
			return &mongo.Database{}, err
		}
		// Check the connection
		helper.LogEvent("INFO", "Confirming MongoDB Connection...")
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			helper.LogEvent("ERROR", helper.ErrorMessage(helper.MONGO_DB_ERROR, err.Error()))
			return &mongo.Database{}, err
		}

		//helper.LogEvent("Info", "Connected to MongoDB!")
		helper.LogEvent("INFO", "Establishing Database collections and indexes...")

		conn = client.Database("transmed")
		return conn, nil

	}

	return conn, nil
}

func CreateIndex(collection *mongo.Collection, field string, unique bool) bool {

	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1}, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(unique),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		helper.LogEvent("ERROR", err.Error())
		return false
	}
	return true
}

func GetPage(page string) (*options.FindOptions, error) {
	if page == "all" {
		return nil, nil
	}
	var limit, e = strconv.ParseInt(helper.Config.PageLimit, 10, 64)
	// var limit, e = strconv.ParseInt(os.Getenv("page_limit"), 10, 64)

	var pageSize, ee = strconv.ParseInt(page, 10, 64)
	if e != nil || ee != nil {
		return nil, helper.ErrorMessage(helper.NO_RECORD_FOUND, "Error in page-size or limit-size.")
	}
	findOptions := options.Find().SetLimit(limit).SetSkip(limit * (pageSize - 1))
	return findOptions, nil
}
