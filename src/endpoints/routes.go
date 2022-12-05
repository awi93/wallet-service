package endpoints

import (
	"net/http"

	"github.com/awi93/wallet-service/src/endpoints/handlers/deposit"
	"github.com/awi93/wallet-service/src/endpoints/handlers/wallet"
	"github.com/awi93/wallet-service/src/utils/httputil"
	"github.com/go-chi/chi"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	dh := deposit.NewHandler(&deposit.Args{
		Service: DepositService,
	})

	r.Post("/deposit", func(w http.ResponseWriter, r *http.Request) {
		httputil.HandleRequest(w, r, dh.Deposit)
	})

	wh := wallet.NewHandler(&wallet.Args{
		Service: WalletService,
	})

	r.Get("/details/{walletId}", func(w http.ResponseWriter, r *http.Request) {
		httputil.HandleRequest(w, r, wh.Wallet)
	})

	return r
}
