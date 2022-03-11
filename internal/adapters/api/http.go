package api

import (
	"encoding/json"
	"wallet/internal/ports"
	"net/http"
)
// bin\colab8
type HTTPHandler struct {
	walletService ports.WalletService
}

func NewHTTPHandler(walletService ports.WalletService, ) *HTTPHandler {
	return &HTTPHandler{
		walletService: walletService,
	}
}

func encodeResult(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(&result)

}
