package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"wallet/internal/adapters/api"
	"wallet/internal/adapters/repositories/mongodb"
	"wallet/internal/adapters/routers"
	"wallet/internal/core/helper"
	"wallet/internal/core/middleware"
	"wallet/internal/core/services"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	helper.InitializeLog()

	db := startMongoDB()

	walletService := walletService(db)
	// mongodb.TotalMoney, _ = mongodb.GetTotalMoney(db)

	handler := api.NewHTTPHandler(walletService)

	if helper.Config.ServiceMode == "dev" {
		router := routers.Router(helper.Config.ServicePort, helper.Config.ServiceAddress, handler)

		helper.LogEvent("Info", fmt.Sprintf("Started HealthServiceApplication on "+helper.Config.ServiceAddress+":"+helper.Config.ServicePort+" in "+time.Since(time.Now()).String()))

		log.Fatal(http.ListenAndServe(":"+helper.Config.ServicePort, middleware.LogRequest(router)))
	}
}

func walletService(db *mongo.Database) *services.WalletService {
	walletRepository := mongodb.NewWalletRepositories(db)
	walletService := services.NewWalletService(walletRepository)
	return walletService
}

func startMongoDB() *mongo.Database {
	helper.LogEvent("INFO", "Initializing Mongo!")
	db, err := mongodb.ConnectToMongo(helper.Config.DbType, helper.Config.MongoDbUserName, helper.Config.MongoDbPassword, helper.Config.MongoDbHost, helper.Config.MongoDbPort, helper.Config.MongoDbAuthDb, helper.Config.MongoDbName, helper.Config.Server)
	if err != nil {
		helper.LogEvent("ERROR", "MongoDB database connection error: "+err.Error())
		log.Fatal()
	}
	return db
}
