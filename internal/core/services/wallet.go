package services

import (
	"wallet/internal/core/domain/entity"
	"wallet/internal/ports"
)

type WalletService struct {
	WalletRepository ports.WalletRepo
}

func NewWalletService(walletRepository ports.WalletRepo) *WalletService {
	return &WalletService{
		WalletRepository: walletRepository,
	}
}

func (serve *WalletService) CreateWallet(wallet entity.Wallet) (interface{}, error) {
	return serve.WalletRepository.CreateWallet(wallet)
}

func (serve *WalletService) CreditWallet(wallet entity.Wallet, totalMoney int, reference string) (interface{}, error) {
	return serve.WalletRepository.CreditWallet(wallet, totalMoney, reference)
}

func (serve *WalletService) DebitWallet(wallet entity.Wallet, totalMoney int, reference string) (interface{}, error) {
	return serve.WalletRepository.DebitWallet(wallet, totalMoney, reference)
}

func (serve *WalletService) GetWallet(reference string) (entity.Wallet, error) {
	return serve.WalletRepository.GetWallet(reference)
}

func (serve *WalletService) GetAllWallet() (int, error) {
	return serve.WalletRepository.GetAllWallet()
}