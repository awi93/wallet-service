package deposit

import (
	"encoding/json"
	"net/http"

	"github.com/awi93/wallet-service/src/dtos"
	depositSvc "github.com/awi93/wallet-service/src/services/deposit"
	"github.com/awi93/wallet-service/src/utils/httputil"
)

type Deposit struct {
	service depositSvc.Service
}

type Args struct {
	Service depositSvc.Service
}

func NewHandler(Args *Args) *Deposit {
	return &Deposit{
		service: Args.Service,
	}
}

func (h *Deposit) Deposit(r *http.Request) httputil.Response {
	var create dtos.DepositRequest
	err := json.NewDecoder(r.Body).Decode(&create)
	if err != nil {
		return httputil.Response{
			StatusCode: 400,
			Body: map[string]interface{}{
				"error": "bad request",
				"cause": "invalid request body from client",
			},
		}
	}

	err = h.service.EmitDeposit(create)
	if err != nil {
		return httputil.Response{
			StatusCode: 500,
			Body: map[string]interface{}{
				"error": "internal server error",
				"cause": "fail to emit deposit event",
			},
		}
	}

	return httputil.Response{
		StatusCode: 201,
		Body: map[string]interface{}{
			"is_success": true,
		},
	}
}
