package wallet

import (
	"net/http"

	"github.com/awi93/wallet-service/src/utils/httputil"
	"github.com/go-chi/chi"
)

func (h *Deposit) Wallet(r *http.Request) httputil.Response {
	walletId := chi.URLParam(r, "walletId")
	wallet, err := h.service.GetWallet(walletId)
	if err != nil {
		if err.Error() == "record not found" {
			return httputil.Response{
				StatusCode: 404,
				Body: map[string]interface{}{
					"error": "not found",
					"cause": "record not found",
				},
			}
		}
		return httputil.Response{
			StatusCode: 500,
			Body: map[string]interface{}{
				"error": "internal server error",
				"cause": "fail to fetch wallet data",
			},
		}
	}
	return httputil.Response{
		StatusCode: 200,
		Body:       wallet,
	}
}
