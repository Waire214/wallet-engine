package ports

import "wallet/internal/core/domain/entity"

type WalletService interface {
	CreateWallet(entity.Wallet) (interface{}, error)
	CreditWallet(entity.Wallet, int, string) (interface{}, error)
	DebitWallet(entity.Wallet, int, string) (interface{}, error)
	GetWallet(string) (entity.Wallet, error)
	GetAllWallet() (int, error)
}
