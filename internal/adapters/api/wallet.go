package api

import (
	"encoding/json"
	"net/http"
	"wallet/internal/core/domain/entity"

	"github.com/go-chi/chi/v5"
)

func (hdl *HTTPHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var wallet entity.Wallet
	err := json.NewDecoder(r.Body).Decode(&wallet)
	if err != nil {
		encodeResult(w, err)
		return
	}
	newWallet, err := hdl.walletService.CreateWallet(wallet)
	if err != nil {
		encodeResult(w, err)
		return
	}
	encodeResult(w, newWallet)
}

// CreateWallet(entity.Wallet) (interface{}, error)
// UpdateWallet(entity.Wallet, string) (interface{}, error)
// GetWallet(string) (entity.Wallet, error)

func (hdl *HTTPHandler) CreditWallet(w http.ResponseWriter, r *http.Request) {
	reference := chi.URLParam(r, "reference")
	var wallet entity.Wallet
	err := json.NewDecoder(r.Body).Decode(&wallet)
	if err != nil {
		encodeResult(w, err)
		return
	}
	totalMoney, _ := hdl.walletService.GetAllWallet()
	creditWallet, err := hdl.walletService.CreditWallet(wallet, totalMoney, reference)
	if err != nil {
		encodeResult(w, err)
		return
	}
	encodeResult(w, creditWallet)
}

func (hdl *HTTPHandler) DebitWallet(w http.ResponseWriter, r *http.Request) {
	reference := chi.URLParam(r, "reference")
	var wallet entity.Wallet
	err := json.NewDecoder(r.Body).Decode(&wallet)
	if err != nil {
		encodeResult(w, err)
		return
	}
	totalMoney, _ := hdl.walletService.GetAllWallet()
	debitWallet, err := hdl.walletService.DebitWallet(wallet, totalMoney, reference)
	if err != nil {
		encodeResult(w, err)
		return
	}
	encodeResult(w, debitWallet)
}

func (hdl *HTTPHandler) GetWallet(w http.ResponseWriter, r *http.Request) {
	reference := chi.URLParam(r, "reference")
	getAWallet, err := hdl.walletService.GetWallet(reference)
	if err != nil {
		encodeResult(w, err)
		return
	}
	encodeResult(w, getAWallet)
}
