package entity

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	Reference         string    `json:"reference" bson:"reference"`
	UserReference     string    `json:"user_reference" bson:"user_reference"`
	Money             int       `json:"money" bson:"money"`
	TotalMoney        int       `json:"total_money" bson:"total_money"`
	CreditedBy        string    `json:"credited_by" bson:"credited_by"`
	CreditorReference string    `json:"creditor_reference" bson:"creditor_reference"`
	CreatedAt         time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at" bson:"deleted_at"`
}

func NewWallet(userReference, creditedBy, creditorReference string, money int) *Wallet {
	var wallet Wallet
	wallet.Reference = uuid.New().String()
	wallet.UserReference = userReference
	wallet.Money = money
	wallet.CreditedBy = creditedBy
	wallet.CreditorReference = creditorReference
	wallet.CreatedAt = time.Now()

	return &wallet
}
