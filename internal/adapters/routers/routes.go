package routers

import (
	// _ "health/docs"
	"wallet/internal/adapters/api"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	// httpSwagger "github.com/swaggo/http-swagger"
)

func Router(appPort, hostAddress string, hdl *api.HTTPHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Mount("/wallet", WalletEndPoints(hdl))
	// router.Get("/swagger/*", httpSwagger.Handler(
	// 	httpSwagger.URL(hostAddress+":"+appPort+"/swagger/doc.json"),
	// ))

	return router
}


func WalletEndPoints(hdl *api.HTTPHandler) http.Handler {
	router := chi.NewRouter()
	router.Post("/", hdl.CreateWallet)
	router.Post("/credit/{reference}", hdl.CreditWallet)
	router.Post("/debit/{reference}", hdl.DebitWallet)
	router.Get("/retrieve/{reference}", hdl.GetWallet)
	return router
}

