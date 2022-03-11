package mongodb

import (
	"context"
	"errors"
	"reflect"
	"time"
	"wallet/internal/core/domain/entity"
	"wallet/internal/core/helper"
	"wallet/internal/ports"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type walletInfra struct {
	walletCollection *mongo.Collection
}

var ctx = context.TODO()
var TotalMoney int

func NewWalletRepositories(conn *mongo.Database) ports.WalletRepo {
	return &walletInfra{
		walletCollection: conn.Collection("wallet_service"),
	}
}

func (conn *walletInfra) CreateWallet(wallet entity.Wallet) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting wallet with reference: "+wallet.Reference)

	newWallet := entity.NewWallet(wallet.UserReference, wallet.CreditedBy, wallet.CreditorReference, wallet.Money)

	_, err := conn.walletCollection.InsertOne(
		ctx,
		newWallet,
	)
	if err != nil {
		return "error", helper.ErrorMessage(helper.MONGO_DB_ERROR, err.Error())
	}

	helper.LogEvent("INFO", "Persisting wallet with reference: "+newWallet.Reference+" completed successfully...")
	return struct {
		Reference string `json:"reference"`
	}{
		Reference: newWallet.Reference,
	}, nil
}

func (conn *walletInfra) CreditWallet(wallet entity.Wallet, totalMoney int, reference string) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting wallet with reference: "+wallet.Reference)

	_, err := conn.walletCollection.InsertOne(
		ctx,
		bson.M{
			"reference":          reference,
			"user_reference":     wallet.UserReference,
			"money":              wallet.Money,
			"total_money":        wallet.Money + totalMoney,
			"credited_by":        wallet.CreditedBy,
			"creditor_reference": wallet.CreditorReference,
			"created_at":         time.Now(),
		},
	)
	if err != nil {
		return "", helper.ErrorMessage(helper.MONGO_DB_ERROR, err.Error())
	}
	helper.LogEvent("INFO", "Persisting wallet with reference: "+reference+" completed successfully...")
	return struct {
		Reference  string `json:"reference"`
		TotalMoney int    `json:"total_money"`
	}{
		Reference:  reference,
		TotalMoney: wallet.Money + totalMoney,
	}, nil

}

func (conn *walletInfra) DebitWallet(wallet entity.Wallet, totalMoney int, reference string) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting wallet with reference: "+wallet.Reference)
	if totalMoney-wallet.Money < 0 {
		return nil, helper.ErrorMessage("Balance_Error", errors.New("insufficient balance").Error())
	}
	_, err := conn.walletCollection.InsertOne(
		ctx,
		bson.M{
			"reference":          uuid.New().String(),
			"user_reference":     wallet.UserReference,
			"money":              wallet.Money,
			"total_money":        totalMoney - wallet.Money,
			"credited_by":        wallet.CreditedBy,
			"creditor_reference": wallet.CreditorReference,
			"created_at":         time.Now(),
		},
	)
	if err != nil {
		return "", helper.ErrorMessage(helper.MONGO_DB_ERROR, err.Error())
	}

	helper.LogEvent("INFO", "Persisting wallet with reference: "+reference+" completed successfully...")
	return struct {
		Reference  string `json:"reference"`
		TotalMoney int    `json:"total_money"`
	}{
		Reference:  reference,
		TotalMoney: totalMoney - wallet.Money,
	}, nil
}

func (conn *walletInfra) GetWallet(reference string) (entity.Wallet, error) {
	helper.LogEvent("INFO", "Retrieving wallet with reference: "+reference)

	var wallet entity.Wallet
	filter := bson.M{"reference": reference}
	err := conn.walletCollection.FindOne(ctx, filter).Decode(&wallet)
	if err != nil || wallet == (entity.Wallet{}) {
		helper.LogEvent("ERROR", helper.NO_RECORD_FOUND)
		return entity.Wallet{}, helper.ErrorMessage(helper.NO_RECORD_FOUND, helper.NO_RECORD_FOUND)
	}

	helper.LogEvent("INFO", "Retrieving wallet with reference: "+reference+" completed successfully...")
	return wallet, nil
}

func (conn *walletInfra) GetAllWallet() (int, error) {
	helper.LogEvent("INFO", "Retrieving all wallet entries...")
	var totalMoney int
	var (
		wallet  entity.Wallet
		wallets []entity.Wallet
	)

	cursor, err := conn.walletCollection.Find(ctx, bson.D{})
	if err != nil {
		return 0, helper.ErrorMessage(helper.NO_RECORD_FOUND, helper.NO_RECORD_FOUND)
	}
	for cursor.Next(ctx) {
		err := cursor.Decode(&wallet)
		if err != nil {
			return 0, helper.ErrorMessage(helper.NO_RECORD_FOUND, err.Error())
		}
		wallets = append(wallets, wallet)
	}
	if reflect.ValueOf(wallets).IsNil() {
		helper.LogEvent("INFO", "There are no results in this collection")
		return 0, nil
	}

	last := wallets[len(wallets)-1]
	totalMoney = last.TotalMoney
	helper.LogEvent("INFO", "Retrieving all wallet entries completed successfully...")
	return totalMoney, nil
}
